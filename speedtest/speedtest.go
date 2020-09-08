package speedtest

import (
	"time"

	st "github.com/kylegrantlucas/speedtest"
)

// Handler run the speedtest repeatly
type Handler interface {
	Run() chan int
}

type handler struct {
	client     Client
	repository Repository
}

// NewHandler creates a new handler for speedtests
func NewHandler(client Client, repository Repository) Handler {
	return handler{client: client, repository: repository}
}

// Run starts a go func with a periodic speedtest
func (h handler) Run() chan int {
	finished := make(chan int)

	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				h.doSpeedtest()
			}
		}
	}()

	return finished
}

func (h handler) doSpeedtest() error {
	test, err := h.client.Test()
	if err != nil {
		return err
	}

	h.repository.Create(test)

	return nil
}

// Client for the speedtest functions
type Client interface {
	Test() (*Speedtest, error)
}

type client struct {
	client *st.Client
}

// NewClient creates new Speedtest Client
func NewClient() (Client, error) {
	c, err := st.NewDefaultClient()
	if err != nil {
		return client{}, err
	}

	return client{client: c}, nil
}

func (c client) Test() (*Speedtest, error) {
	server, err := c.client.GetServer("")
	if err != nil {
		return &Speedtest{}, err
	}

	dl, err := c.client.Download(server)
	if err != nil {
		return &Speedtest{}, err
	}

	ul, err := c.client.Upload(server)
	if err != nil {
		return &Speedtest{}, err
	}

	return NewSpeedtest(dl, ul), nil
}
