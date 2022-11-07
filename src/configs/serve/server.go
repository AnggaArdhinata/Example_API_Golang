package serve

import (
	"log"
	"net/http"
	"os"
	"restapiexample/src/routers"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeComand = &cobra.Command{
	Use: "serve",
	Short: "command to run server",
	RunE: serve,
}

func serve(command *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		c := cors.AllowAll()
		handler := c.Handler(mainRoute)
		var address string = "127.0.0.1:8080"

		if port := os.Getenv("APP_PORT"); port !="" {
			address = "127.0.0.1:" + port
			
		}

		log.Println("Server is running on: " + address)

		if err := http.ListenAndServe(address, handler); err != nil {
			log.Fatal(err.Error())
		}
		return nil
	} else {
		return err
	}
}