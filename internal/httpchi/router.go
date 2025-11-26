package httpchi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/isOdin/RestApi/api/swagger"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type HandlerInterface interface {
	// Auth
	SignUpHandler(w http.ResponseWriter, r *http.Request)
	SignInHandler(w http.ResponseWriter, r *http.Request)

	// Item
	CreateItem(w http.ResponseWriter, r *http.Request)
	GetAllItems(w http.ResponseWriter, r *http.Request)
	GetItemById(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)

	// List
	CreateList(w http.ResponseWriter, r *http.Request)
	GetAllLists(w http.ResponseWriter, r *http.Request)
	GetListById(w http.ResponseWriter, r *http.Request)
	UpdateList(w http.ResponseWriter, r *http.Request)
	DeleteList(w http.ResponseWriter, r *http.Request)
}

type MiddlewareInterface interface {
	JWTAuth(next http.Handler) http.Handler
}

func NewRouter(md MiddlewareInterface, h HandlerInterface) chi.Router {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/swagger/*", httpSwagger.Handler())

	// /auth/...
	r.Route("/api/v0", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-up", h.SignUpHandler)
			r.Post("/sign-in", h.SignInHandler)
		})

		// api/...
		r.Route("/api", func(r chi.Router) {
			r.Use(md.JWTAuth)
			r.Route("/lists", func(r chi.Router) { // api/lists/...
				r.Post("/", h.CreateList)
				r.Get("/", h.GetAllLists)
				r.Get("/{list_id}", h.GetListById)
				r.Put("/{list_id}", h.UpdateList)
				r.Delete("/{list_id}", h.DeleteList)

				r.Route("/{list_id}/items", func(r chi.Router) { // api/lists/items/...
					r.Post("/", h.CreateItem)
				})
				r.Route("/items", func(r chi.Router) { // api/items/...
					r.Get("/", h.GetAllItems)
					r.Get("/{item_id}", h.GetItemById)
					r.Put("/{item_id}", h.UpdateItem)
					r.Delete("/{item_id}", h.DeleteItem)
				})

			})
		})
	})

	return r
}
