package bindchi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/google/uuid"
)

type validator interface {
	Struct(s interface{}) error
}

func DefaultBind[T any](r *http.Request, v *T) error {
	// Создание декодера и создание правило преобразования uuid из string
	decoder := form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return uuid.Parse(vals[0])
	}, uuid.UUID{})

	// ----- Parse request Body -----
	if err := json.NewDecoder(r.Body).Decode(v); err != nil && err != io.EOF {
		return err
	}

	// ----- Parse Url Params -----
	keys := chi.RouteContext(r.Context()).URLParams.Keys
	values := chi.RouteContext(r.Context()).URLParams.Values
	UrlParams := make(map[string][]string)
	for i := 0; i < len(keys); i++ {
		UrlParams[keys[i]] = []string{values[i]}
	}

	// ----- Parse Context -----
	contextValues := make(map[string][]string)

	vValues := reflect.ValueOf(*v)
	elemType := vValues.Type()

	for i := 0; i < vValues.NumField(); i++ {
		// Если поле структуры уже заполнено, то идём на следующую итерацию
		if !vValues.Field(i).IsZero() {
			continue
		}

		// Получаем название поля структурного тэга form
		formTagName := elemType.Field(i).Tag.Get("json")
		if formTagName == "" {
			return errors.New("Not enough structure tags 'json' in structure")
		}

		// Пытаемся получить значение из контекста. Если оно есть, то записываем его в мапу, приводя значение к string
		if ctxValue := r.Context().Value(formTagName); ctxValue != nil {
			contextValues[formTagName] = []string{fmt.Sprintf("%v", ctxValue)}
		}
	}

	maps.Copy(UrlParams, contextValues) // UrlParams and contextValues in one map - in UrlParams

	return decoder.Decode(v, UrlParams)
}

func BindValidate[T any](r *http.Request, v *T, validate validator) error {
	if err := DefaultBind(r, v); err != nil {
		return err
	}

	if err := validate.Struct(v); err != nil {
		return err
	}

	return nil
}
