package handler

import (
	"database/sql"
	"lib_backend/internal/repository"
	"lib_backend/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {

	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)
	loanRepo := repository.NewLoanRepository(db)

	userService := services.NewUserService(userRepo)
	bookService := services.NewBookService(bookRepo)
	loanService := services.NewLoanService(loanRepo, userRepo, bookRepo)

	userHandler := NewUserHandler(userService)
	bookHandler := NewBookHandler(bookService)
	loanHandler := NewLoanHandler(loanService)

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)            // POST /api/users
			users.GET("by-email", userHandler.GetUserByEmail) // GET /api/users/by-email?email=
			users.GET("", userHandler.GetAllUsers)            // GET /api/users
			users.GET(":id", userHandler.GetUserByID)         // GET /api/users/:id
			users.PUT(":id", userHandler.UpdateUser)          // PUT /api/users/:id
			users.DELETE(":id", userHandler.DeleteUser)       // DELETE /api/users/:id
		}

		books := api.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)          // POST /api/books
			books.GET("by-isbn", bookHandler.GetBookByISBN) // GET /api/books/by-isbn?isbn=
			books.GET("", bookHandler.GetAllBooks)          // GET /api/books (deve vir após as rotas mais específicas)
			books.GET(":id", bookHandler.GetBookByID)       // GET /api/books/:id
			books.PUT(":id", bookHandler.UpdateBook)        // PUT /api/books/:id
			books.DELETE(":id", bookHandler.DeleteBook)     // DELETE /api/books/:id
		}

		loans := api.Group("/loans")
		{
			loans.POST("", loanHandler.CreateLoan)          // POST /api/loans
			loans.GET(":id", loanHandler.GetLoanByID)       // GET /api/loans/:id
			loans.PUT(":id/return", loanHandler.ReturnBook) // PUT /api/loans/:id/return
			loans.GET("", loanHandler.GetAllLoans)          // GET /api/loans

			loans.GET("by-user/:user_id", loanHandler.GetLoansByUserID) // GET /api/loans/by-user/:user_id
			loans.GET("by-book/:book_id", loanHandler.GetLoansByBookID) // GET /api/loans/by-book/:book_id

			loans.DELETE(":id", loanHandler.DeleteLoan)
		}
	}
}
