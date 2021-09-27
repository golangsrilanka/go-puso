package puso

import (
	"gorm.io/gorm"

	"github.com/GolangSriLanka/go-puso/internal/maths"
	log "github.com/sirupsen/logrus"

	"github.com/GolangSriLanka/go-puso/models"
)

type PusoRepo interface {
	Save(puso *models.Puso) error
	GetList() ([]models.Puso, error)
	Get(id string) (*models.Puso, error)
	Delete(id string) error
	Update(puso *models.Puso, id string) error
	Migrate() error
}

type Puso struct {
	db *gorm.DB
}

func NewPusoRepo(db *gorm.DB) *Puso {
	return &Puso{
		db: db,
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
