package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/atomic"

	"github.com/goforbroke1006/fake-quotes-svc/domain"
	"github.com/goforbroke1006/fake-quotes-svc/internal/component"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/config"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/shutdowner"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/svc_http"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/wshub"
)

func init() {
	var handleAddrArg = "0.0.0.0:8080"

	var emitWebsocketCmd = &cobra.Command{
		Use:   "websocket",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			var cfg domain.Configuration
			if err := config.ReadFromYaml(&cfg); err != nil {
				panic(err)
			}

			fmt.Println(cfg)

			isReady := atomic.NewBool(false)

			hub := &wshub.WSHub{}

			go func() {
				var wg sync.WaitGroup
				wg.Add(len(cfg.Actives))

				for _, active := range cfg.Actives {
					go func(a domain.Active) {
						time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

						emitter := component.NewEmitter(a, hub, 500)
						go emitter.Emit()

						wg.Done()
					}(active)
				}

				wg.Wait()
				isReady.Store(true)
			}()

			go func() {
				mux := svc_http.New(isReady)

				headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
				originsOk := handlers.AllowedOrigins([]string{"*"})
				methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

				mux.HandleFunc("/ws", handleWsConn(hub))
				logrus.Fatal(http.ListenAndServe(handleAddrArg, handlers.CORS(originsOk, headersOk, methodsOk)(mux)))
			}()

			shutdowner.WaitTerminateSignal(func() {
				_ = hub.Close()
			})
		},
	}

	emitCmd.Flags().StringVarP(&handleAddrArg, "handle-addr", "a", handleAddrArg, "HTTP server endpoint")
	emitCmd.AddCommand(emitWebsocketCmd)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWsConn(hub *wshub.WSHub) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		log.Print("new connection:")
		hub.Add(c)
	}
}
