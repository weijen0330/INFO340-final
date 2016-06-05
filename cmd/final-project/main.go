// by setting package as main, Go will compile this as an executable file.
// Any other package turns this into a library
package main

// These are your imports / libraries / frameworks
import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	// The pointer to the database 
	db *sql.DB

)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var errd error
	// Open a connection to the database using an environemnt variable.
	// Not the best technique, but the simplest one for heroku
	db, errd = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errd != nil {
		log.Fatalf("Error opening database: %q", errd)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("html/*")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/ping", func(c *gin.Context) {
		ping := db.Ping()
		if ping != nil {
			// our site can't handle http status codes, but I'll still put them in cause why not
			c.JSON(http.StatusOK, gin.H{"error": "true", "message": "db was not created. Check your DATABASE_URL"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "false", "message": "db created"})
		}
	})

	router.GET("/query1", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"
		// put your query here
		rows, err := db.Query("SELECT * from recentTrips;")
		if err != nil {
			// careful about returning errors to the user!
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		// foreach loop over rows.Columns, using value
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all returned columns here  
		var ID int  
		var date string
		var destination	string
		var origin string 

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&ID, &date, &destination, &origin) 
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>" + strconv.Itoa(ID) + "</td><td>" + date[:10]+ "</td><td>" + destination + "</td><td>" + origin + "</td></tr>"
		}
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/update", func(c *gin.Context) {
		table := "<table class='table'><thead><tr>"

		searchBox := c.PostForm("searchBox")
		likeString := searchBox + "%"
		// "SELECT * FROM recentTrips WHERE  LIKE $2", input, likeString
		rows, err := db.Query("SELECT * FROM recentTrips WHERE destination LIKE $1;", likeString)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		for _, value := range cols {
			table += "<th class='text-center'>" + value + "</th>"
		}
		// once added all the columns in, close the header
		table += "</thead><tbody>"
		// declare all returned columns here  
		var ID int  
		var date string
		var destination	string
		var origin string 

		for rows.Next() {
			// assign each of them, in order, to the parameters of rows.Scan.
			// preface each variable with &
			rows.Scan(&ID, &date, &destination, &origin) 
			// can't combine ints and strings in Go. Use strconv.Itoa(int) instead
			table += "<tr><td>" + strconv.Itoa(ID) + "</td><td>" + date[:10]+ "</td><td>" + destination + "</td><td>" + origin + "</td></tr>"
		}
		c.Data(http.StatusOK, "text/html", []byte(table))
	})

	router.POST("/insert", func(c *gin.Context) {
		
		firstName := c.PostForm("firstName")
		middleName := c.PostForm("middleName")
		lastName := c.PostForm("lastName")
		description := c.PostForm("description") 

		rows, err := db.Query("SELECT addNewCargo($1, $2, $3, $4)", firstName, middleName, lastName, description)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		cols, _ := rows.Columns()
		if len(cols) == 0 {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		var resultUser string
		for rows.Next() {
			rows.Scan(&resultUser)
		} 
		c.JSON(http.StatusOK, gin.H{"firstName": firstName})
	})
	// NO code should go after this line. It won't ever reach that point.
	router.Run(":" + port)
} 
