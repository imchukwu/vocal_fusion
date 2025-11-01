package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"vocal_fusion/internals/handlers"
	// "vocal_fusion/internals/middleware"
	"vocal_fusion/internals/repository"
)

func RegisterRoutes(r *chi.Mux, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🎤 Welcome to Vocal Fusion API"))
	})

	r.Route("/users", func(api chi.Router) {
		api.Post("/register", userHandler.RegisterUser)
		api.Post("/login", userHandler.Login)
		api.Get("/", userHandler.GetAllUsers)
		api.Get("/{id}", userHandler.GetUserByID)
		api.Put("/{id}", userHandler.UpdateUser)
		api.Delete("/{id}", userHandler.DeleteUser)
	})

	// ===== Events =====
	eventRepo := repository.NewEventRepository(db)
	eventHandler := handlers.NewEventHandler(eventRepo)

	r.Route("/events", func(api chi.Router) {
		api.Post("/", eventHandler.CreateEvent)
		api.Get("/", eventHandler.GetEvents)
		api.Get("/{id}", eventHandler.GetEventByID)
		api.Put("/{id}", eventHandler.UpdateEvent)
		api.Delete("/{id}", eventHandler.DeleteEvent)
	})

	// ===== FAQs =====
	faqRepo := repository.NewFAQRepository(db)
	faqHandler := handlers.NewFAQHandler(faqRepo)

	r.Route("/faqs", func(api chi.Router) {
		api.Post("/", faqHandler.CreateFAQ)
		api.Get("/", faqHandler.GetAllFAQs)
		api.Get("/{id}", faqHandler.GetFAQByID)
		api.Put("/{id}", faqHandler.UpdateFAQ)
		api.Delete("/{id}", faqHandler.DeleteFAQ)
	})

	// ===== Media =====
	mediaRepo := repository.NewMediaRepository(db)
	mediaHandler := handlers.NewMediaHandler(mediaRepo)

	r.Route("/media", func(api chi.Router) {
		api.Post("/", mediaHandler.CreateMedia)
		api.Get("/", mediaHandler.GetAllMedia)
		api.Get("/{id}", mediaHandler.GetMediaByID)
		api.Put("/{id}", mediaHandler.UpdateMedia)
		api.Delete("/{id}", mediaHandler.DeleteMedia)
	})

	messageRepo := repository.NewMessageRepository(db)
	messageHandler := handlers.NewMessageHandler(messageRepo)

	r.Route("/messages", func(api chi.Router) {
		api.Post("/", messageHandler.CreateMessage)
		api.Get("/", messageHandler.GetAllMessages)
		api.Get("/{id}", messageHandler.GetMessageByID)
		api.Patch("/{id}/status", messageHandler.UpdateMessageStatus)
		api.Delete("/{id}", messageHandler.DeleteMessage)
	})

	schoolRepo := repository.NewSchoolRepository(db)
	schoolHandler := handlers.NewSchoolHandler(schoolRepo)

	r.Route("/schools", func(api chi.Router) {
		api.Post("/", schoolHandler.CreateSchool)
		api.Get("/", schoolHandler.GetAllSchools)
		api.Get("/{id}", schoolHandler.GetSchoolByID)
		api.Put("/{id}", schoolHandler.UpdateSchool)
		api.Delete("/{id}", schoolHandler.DeleteSchool)
	})

	// WinnerSays routes
	winnerRepo := repository.NewWinnerSaysRepository(db)
	winnerHandler := handlers.NewWinnerSaysHandler(winnerRepo)

	r.Route("/winnersays", func(r chi.Router) {
		r.Post("/", winnerHandler.CreateWinnerSays)
		r.Get("/", winnerHandler.GetAllWinnerSays)
		r.Get("/{id}", winnerHandler.GetWinnerSaysByID)
		r.Delete("/{id}", winnerHandler.DeleteWinnerSays)
	})

}
