package bootstrap

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 应用配置
type Configurator func(bootstrapper *Bootstrapper)

type Bootstrapper struct {
	*gin.Engine
	AppName     string
	AppOwner    string
	AppSpawDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Engine:      gin.Default(),
		AppName:     appName,
		AppOwner:    appOwner,
		AppSpawDate: time.Now(),
	}
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

// 其他设置，比如view，static，errorHandler todo。。。

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) SetupCron() {
	// todo
}

const (
	StaticAsset = "public/"
	Favicon     = "favicon.ico"
)

func (b *Bootstrapper) BootStrap() *Bootstrapper {
	b.Use(gin.Logger())
	b.Use(gin.Recovery())
	// some setup todo
	b.SetupCron()

	b.Static("/assets", StaticAsset)
	b.StaticFile("favicon.ico", StaticAsset+Favicon)
	return b
}

func (b *Bootstrapper) Start(addr string) {
	s := &http.Server{
		Addr:           addr,
		Handler:        b,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
