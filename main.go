package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	db1 "server/db1"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	toml "github.com/pelletier/go-toml"
)

type (
	User struct {
		UserId   int    `json:"userid"`
		UserName string `json:"username"`
		Pass     string `json:"pass"`
		//Authors []Author `json:"authors"`
	}
)

type (
	Users struct {
		Users []User `json:"users"`
	}
)

type (
	Authors struct {
		Authors []Author `json:"authors"`
	}
)

type (
	ListName struct {
		UId   int    `json:"userid"`
		UName string `json:"username"`
		UPass string `json:"pass"`
	}
)

/*
type (
	response struct {
		Success bool   `json:"success"`
		Message string `json:"messsage,omitempty"`
	}
)
*/

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db1"
)

//for public schema author table
/*
const (
	host     = "gononeterp.cud7jbsftjfi.ap-southeast-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "z-pMZ2e?+K"
	dbname   = "choukash_erp_1_0_0"
)
*/

func dbconnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func goUser(c echo.Context) error {
	return c.String(http.StatusOK, "this is for goUser!")
}
func pythonUser(c echo.Context) error {
	return c.String(http.StatusOK, "this is for python user!")
}

func javaUser(c echo.Context) error {
	return c.String(http.StatusOK, "this is for javaUser only!")
}

func createAuthor(c echo.Context) error {
	db := dbconnection()
	//u := &User{}
	a := &db1.Author{}
	if err := c.Bind(a); err != nil {
		return err
	}
	/*
		fmt.Println(u.Name)
		fmt.Println(u.Pass)
		var name = "kabir"
		var pass = "123"
		r := response{Success: true}
		if name != u.Name || pass != u.Pass {
			r.Success = false
			r.Message = "Authentication failed"
			return c.JSON(http.StatusOK, r)
		}
	*/

	//sqlStatemnt := "INSERT INTO authors(authorid, authorname, authorbio)VALUES ($1, $2, $3)"

	dtbase := db1.New(db)
	author, err := dtbase.CreateAuthor(context.Background(), db1.CreateAuthorParams{
		//Name: "jahir",
		//Bio:  "dfgdfgh",
		Name: a.Name,
		Bio:  a.Bio,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(author)
	/*
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, a)
		}
	*/
	return c.String(http.StatusOK, "ok")

}

//updateAuthor

func updateAuthor(c echo.Context) error {
	db := dbconnection()
	a := &db1.Author{}
	if err := c.Bind(a); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	Authors[id].Name = a.Name
	//fmt.Print(author)
	dtbase := db1.New(db)
	var uauthor, err = dtbase.UpdateAuthor(context.Background(), db1.UpdateAuthorParams{
		//Name: a.Name,
		Bio: a.Bio,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	return c.JSON(http.StatusOK, uauthor)
	/*
			func updateUser(c echo.Context) error {
			u := new(user)
			if err := c.Bind(u); err != nil {
				return err
			}
			id, _ := strconv.Atoi(c.Param("id"))
			users[id].Name = u.Name
			return c.JSON(http.StatusOK, users[id])
		}
	*/
}

//deleteAuthor

func deleteAuthor(c echo.Context) error {
	db := dbconnection()
	//a := &db1.Author{}
	author, _ := strconv.Atoi(c.Param("id"))

	dtbase := db1.New(db)
	dauthor := dtbase.DeleteAuthor(context.Background(), int64(author))
	return c.JSON(http.StatusOK, dauthor)
}

/*
func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
*/

//getAuthor

func getAllAuthor(c echo.Context) error {
	db := dbconnection()
	dtbase := db1.New(db)
	author, err := dtbase.GetAuthor(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(author)
	//parchaseItem, err := dtbase.
	//	q := &hello.Queries{}

	//db := db1.New(dtbase)

	/*names, err := db1.(context.Background())
	if err != nil {
		return err
	}*/
	//	fmt.Println(names)

	//	return nil
	return c.JSON(http.StatusOK, author)

}

/*
	fmt.Println(u.Name)
	fmt.Println(u.Pass)
	var name = "kabir"
	var pass = "123"
	r := response{Success: true}
	if name != u.Name || pass != u.Pass {
		r.Success = false
		r.Message = "Authentication failed"
		return c.JSON(http.StatusOK, r)
	}
*/

func createUser(c echo.Context) error {

	db := dbconnection()
	u := &db1.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	dtbase := db1.New(db)
	//names, _ := db.ListAll(context.Background())

	user, err := dtbase.CreateUser(context.Background(), db1.CreateUserParams{
		Username: u.Username,
		Pass:     u.Pass,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

	return c.String(http.StatusOK, "ok")
}

func updateUser(c echo.Context) error {

	//db := dbconnection()

	u := &db1.Author{}
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.String(http.StatusOK, u.Name)
}

func deleteUser(c echo.Context) error {
	db := dbconnection()
	userid := c.Param("userid")
	sqlStatement := "DELETE FROM users WHERE authorid = $1"
	res, err := db.Query(sqlStatement, userid)
	if err != nil {
		fmt.Println(err)
		//return c.JSON(http.StatusCreated, u);
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusOK, "Deleted")
	}
	return c.String(http.StatusOK, userid+" Deleted")
}

func getAllUser(c echo.Context) error {
	db := dbconnection()
	dtbase := db1.New(db)
	user, err := dtbase.GetUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)

	return c.JSON(http.StatusOK, user)

}

/*
func getUserId(c echo.Context) error {
	id := c.Param("id")

	if id == "1" {
		return c.String(http.StatusOK, "hi this is joe")
	} else {

		return c.String(http.StatusOK, "hi this is hello")
	}
}
*/
/*
func getProduct(c echo.Context) error {
	return c.String(http.StatusOK, "here is your product list!")
}
*/

func main() {

	config, _ := toml.LoadFile("config.ini")

	dbconfig := config.Get("database").(*toml.Tree)

	db := dbconn(dbconfig)
	defer db.Close()

	//database := db1.New(db)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//e.Logger.Fatal(e.Start(":1323"))
	e.GET("/goUser", goUser)
	e.GET("/pythonUser", pythonUser)
	e.GET("/javaUser/", javaUser)

	//path matching order + user id path binding
	//e.GET("/users/:id", getUserId)
	//e.GET("/products", getProduct)

	//autor operation
	e.POST("/createAuthor", createAuthor)
	e.PUT("/updateAuthor/:id", updateAuthor)
	e.DELETE("/deleteAuthor/:id", deleteAuthor)
	e.GET("/getAllAuthor", getAllAuthor)

	//user operation
	e.POST("/createUser", createUser)
	e.PUT("/updateUser/:id", updateUser)
	e.DELETE("/deleteUser", deleteUser)
	e.GET("/getAllUser", getAllUser)

	e.Logger.Fatal(e.Start(":8080"))

}
