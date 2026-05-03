package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"vocal_fusion/internals/handlers"
	"vocal_fusion/internals/middleware"
	"vocal_fusion/internals/repository"
	"vocal_fusion/pkg/email"
)

func RegisterRoutes(r *chi.Mux, db *gorm.DB, emailSvc email.EmailService) {
	// Global Rate Limiting (5 requests per second, burst of 10)
	rl := middleware.NewRateLimiter(5, 10)
	r.Use(rl.Limit)

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🎤 Welcome to Vocal Fusion API"))
	})

	r.Route("/users", func(api chi.Router) {
		api.Post("/register", userHandler.RegisterUser)
		api.Post("/login", userHandler.Login)
		
		// Protected
		api.Group(func(protected chi.Router) {
			protected.Use(middleware.JWTAuthMiddleware)
			protected.Get("/", userHandler.GetAllUsers)
			protected.Get("/{id}", userHandler.GetUserByID)
			protected.Put("/{id}", userHandler.UpdateUser)
			protected.Delete("/{id}", userHandler.DeleteUser)
		})
	})

	// ===== Events =====
	eventRepo := repository.NewEventRepository(db)
	eventHandler := handlers.NewEventHandler(eventRepo)

	r.Route("/events", func(api chi.Router) {
		api.Post("/", eventHandler.CreateEvent)
		api.Get("/", eventHandler.GetEvents)
		api.Get("/types", eventHandler.GetEventTypes) // ✅ Added event types route
		api.Get("/count", eventHandler.GetEventCount) // ✅ Added count route
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
	messageHandler := handlers.NewMessageHandler(messageRepo, emailSvc)

	r.Route("/messages", func(api chi.Router) {
		api.Post("/", messageHandler.CreateMessage)
		api.Post("/bulk", messageHandler.SendBulkMessage)
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
		api.Patch("/{id}/confirm", schoolHandler.UpdateConfirmationStatus) // ✅ Added confirm route
		api.Delete("/{id}", schoolHandler.DeleteSchool)
	})

	schoolEventRepo := repository.NewSchoolEventRepository(db)
	schoolEventHandler := handlers.NewSchoolEventHandler(schoolEventRepo)

	r.Route("/registrations", func(api chi.Router) {
		api.Get("/", schoolEventHandler.GetAllRegistrations)
		api.Get("/events", schoolEventHandler.GetAllRegistrations)
		api.Post("/events/{eventID}", schoolEventHandler.RegisterSchool)
		api.Get("/events/{eventID}", schoolEventHandler.GetEventRegistrations)
		api.Get("/schools/{schoolID}", schoolEventHandler.GetSchoolRegistrations)
		api.Patch("/events/{eventID}/schools/{schoolID}/verify", schoolEventHandler.VerifyRegistration)
		api.Put("/events/{eventID}/schools/{schoolID}/generate-code", schoolEventHandler.GenerateSchoolEventCode)
		api.Put("/events/{eventID}/schools/{schoolID}/code", schoolEventHandler.UpdateSchoolEventCode)
		api.Delete("/events/{eventID}/schools/{schoolID}", schoolEventHandler.UnregisterSchool)
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

	// Settings routes
	settingsRepo := repository.NewSettingsRepository(db)
	settingsHandler := handlers.NewSettingsHandler(settingsRepo)

	r.Route("/settings", func(r chi.Router) {
		r.Get("/", settingsHandler.GetSettings)
		r.Put("/", settingsHandler.UpdateSettings)
	})

	// ===== File Uploads =====
	r.Post("/upload", handlers.UploadFile)

	// Serve static files from the "uploads" directory
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
}
