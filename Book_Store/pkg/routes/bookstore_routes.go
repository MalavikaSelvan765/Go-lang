package routes

import(
    "github.com/gorilla/mux"
    "github.com/MalavikaSelvan765/Go-lang/Book_Store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
    router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
    router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("POST")
    router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}