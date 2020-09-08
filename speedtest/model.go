package speedtest

import (
	"gorm.io/gorm"

	"github.com/spie/gospeed/db"
)

// Speedtest Model
type Speedtest struct {
	gorm.Model
	Download float64
	Upload   float64
}

// NewSpeedtest creates new Speedtest Model
func NewSpeedtest(download float64, upload float64) *Speedtest {
	return &Speedtest{Download: download, Upload: upload}
}

// Repository for the Speedtest model
type Repository interface {
	AutoMigrate()
	Create(speedtest *Speedtest)
	FindAll() []Speedtest
}

type repository struct {
	connectionHandler db.ConnectionHandler
}

// NewRepository creates new Speedtes Repository
func NewRepository(connectionHandler db.ConnectionHandler) Repository {
	return &repository{connectionHandler: connectionHandler}
}

func (r repository) AutoMigrate() {
	r.connectionHandler.AutoMigrate(&Speedtest{})
}

func (r repository) Create(speedtest *Speedtest) {
	r.connectionHandler.Create(speedtest)
}

func (r repository) FindAll() []Speedtest {
	var speedtests []Speedtest

	r.connectionHandler.Find(&Speedtest{})

	return speedtests
}
