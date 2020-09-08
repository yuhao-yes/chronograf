package dbstore

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/influxdata/chronograf"
	"github.com/influxdata/chronograf/id"
	"github.com/influxdata/chronograf/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Ensure dashboardsStore implements chronograf.DashboardsStore.
var _ chronograf.DashboardsStore = &DashboardsStore{}

var (
	generator *snowflake.Generator
	uuid      *id.UUID
)

func init() {
	generator = snowflake.New(rand.Intn(1023))
	uuid = &id.UUID{}
}

// DashboardsStore implements chronograf.DashboardsStore
type DashboardsStore struct {
	db *gorm.DB
}

// NewDashboardsStore create and open dashboard store
func NewDashboardsStore() *DashboardsStore {
	dsn := "root:123456@tcp(localhost:3306)/gxFastDashboards"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&dashboardData{})

	return &DashboardsStore{db: db}
}

// All lists all dashboards from the DashboardStore
func (s *DashboardsStore) All(context.Context) ([]chronograf.Dashboard, error) {
	var datas []dashboardData
	s.db.Find(&datas)

	result := make([]chronograf.Dashboard, len(datas))
	for i, d := range datas {
		d.unload(&result[i])
	}
	return result, nil
}

// Add Create a new Dashboard in the DashboardStore
func (s *DashboardsStore) Add(ctx context.Context, d chronograf.Dashboard) (chronograf.Dashboard, error) {
	// update id for dashboard and cells
	d.ID = chronograf.DashboardID(nextID())
	for i, cell := range d.Cells {
		cell.ID = nextUUID()
		d.Cells[i] = cell
	}

	//create a data object to save
	data := dashboardData{}
	data.load(&d)

	s.db.Create(&data)
	return d, nil
}

// Delete the Dashboard from the DashboardStore if `ID` exists.
func (s *DashboardsStore) Delete(ctx context.Context, d chronograf.Dashboard) error {
	s.db.Delete(&dashboardData{}, d.ID)
	return nil
}

// Get retrieves a dashboard if `ID` exists.
func (s *DashboardsStore) Get(ctx context.Context, id chronograf.DashboardID) (chronograf.Dashboard, error) {
	data := dashboardData{}
	s.db.First(&data, id)
	result := chronograf.Dashboard{}
	data.unload(&result)
	return result, nil
}

// Update replaces the dashboard information
func (s *DashboardsStore) Update(ctx context.Context, d chronograf.Dashboard) error {
	data := dashboardData{}
	data.load(&d)
	s.db.Save(&data)
	return nil
}

// dashboardData is the object to save by gorm
type dashboardData struct {
	ID   int
	JSON string
}

func (d *dashboardData) unload(target *chronograf.Dashboard) {
	json.Unmarshal([]byte(d.JSON), target)
}

func (d *dashboardData) load(src *chronograf.Dashboard) {
	b, _ := json.Marshal(src)
	d.ID = int(src.ID)
	d.JSON = string(b)
}

func nextID() uint64 {
	return generator.Next()
}

func nextUUID() string {
	result, _ := uuid.Generate()
	return result
}
