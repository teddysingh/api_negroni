# Building a basic REST API with Negroni + GORM 

**HWService**

Hello World Service is a basic service that shows how you can bring up a basic REST API server using plain `net/http` + `negroni`.

---

**User Service**

This implementation shows how you can make use of `gorm` + `negroni` + `net/http` package to bring up a REST API Server that can 

- Create User
- Search/Find a User by `email`
- I have taken a step further to introduce some utils to pre-format the API Response specifying content type or a specific HTTP Status code in case of error.

You should be able to bring up the server using `go run server.go` assuming that the `db.go` file has correct DB connectivity parameters.

Following 3 APIs are exposed in UserService

    GET http://localhost:3000/
    
    Response
    {
    "name": "Hello World Server",
    "port": ":3000",
    "version": "1.0"
    }

    POST http://localhost:3000/users
    {
	"name": "Vidit",
	"email": "vidit@austinmonks.com",
	"password": "1234"
	}
	
    Response
    {
    "ID": 2,
    "CreatedAt": "2018-09-19T15:30:19.165834+05:30",
    "UpdatedAt": "2018-09-19T15:30:19.165834+05:30",
    "DeletedAt": null,
    "Name": "Vidit",
    "Email": "vidit@austinmonks.com",
    "Password": "1234"
	}

    GET http://localhost:3000/users?email=au
    
    Response
    {
        "ID": 2,
        "CreatedAt": "2018-09-19T15:30:19.165834+05:30",
        "UpdatedAt": "2018-09-19T15:30:19.165834+05:30",
        "DeletedAt": null,
        "Name": "Vidit",
        "Email": "vidit@austinmonks.com",
        "Password": "1234"
    }
    

This project is just a PoC to play around with the above mentioned package and of course needs much more polishing before being a production ready code. :)

---

**What next?**

Will be playing next with `FastHTTP` and `gin-gonic` soon.
