package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/jake-dog/opensimdash/codemasters"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Config struct
type Config struct {
	Port      int `json:"port"`
	Lock2Lock int `json:"lock2lock"`
	WindowX   int `json:"window_x"`
	WindowY   int `json:"window_y"`
	WindowW   int `json:"window_w"`
	WindowH   int `json:"window_h"`
}

// App struct
type App struct {
	ctx    context.Context
	show   chan bool
	config *Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{show: make(chan bool)}
}

func (a *App) Kick() {
	a.show <- true
}

func (a *App) handle(ctx context.Context, c *net.UDPConn) {
	b := make([]byte, 4096)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		n, err := c.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		packet := b[:n]
		var pkt codemasters.DirtPacket
		pkt.Decode(packet)
		runtime.EventsEmit(ctx, "telemetry", pkt, a.config.Lock2Lock)
		a.Kick()
		log.Printf("%#v", pkt)
	}
}

func (a *App) configPath() string {
	self, err := os.Executable()
	if err != nil {
		return "params.json"
	}
	return filepath.Join(filepath.Dir(self), "params.json")
}

func (a *App) Load(ctx context.Context) (*Config, error) {
	fp, err := os.Open(a.configPath())
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var conf *Config
	if err := json.NewDecoder(fp).Decode(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func (a *App) Save(ctx context.Context) (err error) {
	conf, err := a.Load(ctx)
	if err != nil {
		conf = &Config{
			Port:      20777,
			Lock2Lock: 900,
			WindowX:   880,
			WindowY:   540,
			WindowW:   230,
			WindowH:   120,
		}
	} else {
		if err := os.Rename(a.configPath(), a.configPath()+".bak"); err != nil {
			return err
		}
	}
	conf.WindowX, conf.WindowY = runtime.WindowGetPosition(ctx)
	conf.WindowW, conf.WindowH = runtime.WindowGetSize(ctx)
	fp, err := os.Create(a.configPath())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if _, e := os.Stat(a.configPath()); os.IsExist(e) {
				os.Remove(a.configPath())
			}
			if _, e := os.Stat(a.configPath() + ".bak"); os.IsExist(e) {
				os.Rename(a.configPath()+".bak", a.configPath())
			}
		}
	}()
	defer fp.Close()
	b, err := json.MarshalIndent(conf, "", "  ")
	if _, err := fp.Write(b); err != nil {
		return err
	}
	if err := fp.Sync(); err != nil {
		return err
	}
	return nil
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	conf, err := a.Load(ctx)
	if err != nil {
		log.Print(err)
		return
	}
	a.config = conf
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("localhost"),
		Port: a.config.Port,
	}
	c, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	runtime.EventsOn(ctx, "window-activate", func(optionalData ...interface{}) {
		a.Kick()
	})
	go a.handle(ctx, c)
	go func() {
		tm := time.NewTimer(15 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-a.show:
				if v {
					tm.Reset(15 * time.Second)
					runtime.WindowShow(ctx)
				}
			case <-tm.C:
				if err := a.Save(ctx); err != nil {
					log.Print(err)
				}
				runtime.WindowHide(ctx)
			}
		}
	}()
	runtime.WindowSetPosition(ctx, a.config.WindowX, a.config.WindowY)
	runtime.WindowSetSize(ctx, a.config.WindowW, a.config.WindowH)
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
