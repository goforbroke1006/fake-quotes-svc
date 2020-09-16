package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/atomic"

	"github.com/goforbroke1006/fake-quotes-svc/domain"
	"github.com/goforbroke1006/fake-quotes-svc/internal/component/wshub"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/config"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/shutdowner"
	"github.com/goforbroke1006/fake-quotes-svc/pkg/svc_http"
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

			stream := make(chan domain.Quote, len(cfg.Actives))
			hub := wshub.WSHub{}

			go func() {
				for m := range stream {
					hub.Send(m)
				}
			}()

			go func() {
				var wg sync.WaitGroup
				wg.Add(len(cfg.Actives))

				for _, active := range cfg.Actives {
					go func(a domain.Active) {
						time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
						go func() {
							for {
								stream <- domain.Quote{
									Code: a.Code,
									Bid:  1,
									Ask:  1,
									At:   time.Now().Unix(),
								} // TODO: realize quote value rand modification
								time.Sleep(100 * time.Millisecond)
							}
						}()
						wg.Done()
					}(active)
				}

				wg.Wait()
				isReady.Store(true)
			}()

			go func() {
				mux := svc_http.New(isReady)
				mux.HandleFunc("/ws", handleWsConn(&hub))
				logrus.Fatal(http.ListenAndServe(handleAddrArg, mux))
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
		hub.Add(c)
	}
}
