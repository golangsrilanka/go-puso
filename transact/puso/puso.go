package puso

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/golangsrilanka/go-puso/internal/maths"
	"github.com/golangsrilanka/go-puso/models"
)

type PusoRepo interface {
	Save(puso *models.Puso) error
	GetList() ([]models.Puso, error)
	Get(id string) (*models.Puso, error)
	Delete(id string) error
	Update(puso *models.Puso, id string) error
	Migrate() error
}

type Config struct {
	fx.In

	DB *gorm.DB
}

type Puso struct {
	db *gorm.DB
}

func NewPusoRepo(config Config) *Puso {
	return &Puso{
		db: config.DB,
	}
}

func (p *Puso) Save(puso *models.Puso) error {
	puso.Laziness = maths.LazinessGen(puso.Name, puso.Color, puso.Weight)
	if result := p.db.Create(&puso); result.Error != nil {
		log.Errorf("failed to create puso: %+v: %v", p, result.Error)
		return result.Error
	}

	return nil
}

func (p *Puso) GetList() ([]models.Puso, error) {
	list := make([]models.Puso, 0)

	if result := p.db.Find(&list); result.Error != nil {
		log.Errorf("failed to find puso: %+v: %v", p, result.Error)
		return nil, result.Error
	}

	return list, nil
}

func (p *Puso) Get(id string) (*models.Puso, error) {
	t := &models.Puso{}

	if result := p.db.Where("ID = ?", id).First(t); result.Error != nil {
		log.Errorf("failed to find puso: %+v: %v", p, result.Error)
		return t, result.Error
	}

	return t, nil
}

func (p *Puso) Delete(id string) error {
	if result := p.db.Model(&p).Where("ID = ?", id).Delete(p); result.Error != nil {
		log.Errorf("failed to delete puso: %+v: %v", p, result.Error)
		return result.Error
	}

	return nil
}

func (p *Puso) Update(puso *models.Puso, id string) error {
	if result := p.db.Model(&puso).Where("ID = ?", id).Updates(puso); result.Error != nil {
		log.Errorf("failed to update puso: %+v: %v", p, result.Error)
		return result.Error
	}

	return nil
}

func (p *Puso) Migrate() error {
	if err := p.db.AutoMigrate(models.Puso{}); err != nil {
		log.Errorf("failed to migrate puso: %+v: %v", models.Puso{}, err)
		return err
	}

	return nil
}
