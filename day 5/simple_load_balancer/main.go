package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive() (alive bool) {
	b.mux.Lock()
	alive = b.Alive
	b.mux.Unlock()
	return
}

func (s *ServerPool) GetNextIndex() int {
	return int(atomic.AddUint64(&s.current, 1)) % len(s.backends)
}

func (s *ServerPool) AddBackend(backend *Backend) {
	s.backends = append(s.backends, backend)
}

func (s *ServerPool) GetNextPeer() *Backend {
	nextIndex := s.GetNextIndex()
	l := len(s.backends) + nextIndex
	for i := nextIndex; i < l; i++ {
		idx := i % len(s.backends)
		if s.backends[idx].IsAlive() {
			if i != nextIndex {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	return nil
}

func loadBalancer(w http.ResponseWriter, r *http.Request) {
	peer := serverPool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
	}
	http.Error(w, "Service not available", http.StatusServiceUnavailable)
}

func isBackendAlive(u *url.URL) bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", u.Host, timeout)
	if err != nil {
		log.Println("Site unreachable, error: ", err)
		return false
	}
	_ = conn.Close()
	return true
}

func (s *ServerPool) HealthCheck() {
	for _, v := range s.backends {
		status := "up"
		alive := isBackendAlive(v.URL)
		v.SetAlive(alive)
		if !alive {
			status = "down"
		}
		log.Printf("%s [%s]\n", v.URL, status)
	}
}

func healthCheck() {
	t := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-t.C:
			log.Println("Starting health check...")
			serverPool.HealthCheck()
			log.Println("Health check completed")
		}
	}
}

var serverPool ServerPool

func main() {
	serverList := []string{"http://127.0.0.1:8100", "http://127.0.0.1:8200"}
	for _, v := range serverList {
		serverURL, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(serverURL)
		serverPool.AddBackend(&Backend{
			URL:          serverURL,
			Alive:        true,
			ReverseProxy: proxy,
		})
	}
	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: http.HandlerFunc(loadBalancer),
	}
	go healthCheck()
	log.Printf("🚀 Load balancer started at 127.0.0.1:8000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
