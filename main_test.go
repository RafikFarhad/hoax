package main_test

import (
	"flag"
	"github.com/RafikFarhad/hoax/app"
	"github.com/RafikFarhad/hoax/config"
	http2 "github.com/RafikFarhad/hoax/http"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

var hostAddress = flag.String("h", "127.0.0.1:3333", "host address")
var configFile = flag.String("c", "test.config.ini", "config file")

func TestMain(m *testing.M) {
	flag.Parse()
	// parse .ini
	if err := config.ParseConfig(*hostAddress, *configFile); err != nil {
		panic(err)
	}
	// create app
	if err := app.CreateApp(true); err != nil {
		panic(err)
	}
	go func() {
		if err := http2.AppHttp.Listen(config.AppConfig.Address); err != nil {
			panic(err)
		}
	}()
	time.Sleep(2 * time.Second)
	st := m.Run()
	os.Exit(st)
}

func TestAppIsListening(t *testing.T) {
	if err := doRequest("http://" + *hostAddress); err != nil {
		t.Error(err)
	}
}

func doRequest(uri string) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
