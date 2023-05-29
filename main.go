package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	albums "github.com/andeemarks/golang-gin/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/albums/:id", getAlbum)

	// The remaining code is taken from
	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	handleShutdown(*server)
}

func handleShutdown(server http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

/*
curl http://localhost:8080/albums \
--header "Content-Type: application/json" \
--request "GET"
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums.GetAll())
}

/*
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
*/
func addAlbum(c *gin.Context) {
	var newAlbum albums.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	c.IndentedJSON(http.StatusCreated, albums.Add(newAlbum))
}

/*
 curl http://localhost:8080/albums/2
*/
func getAlbum(c *gin.Context) {
	foundAlbum, err := albums.Get(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	c.IndentedJSON(http.StatusOK, foundAlbum)
}
