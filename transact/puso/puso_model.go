package puso

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/GolangSriLanka/go-puso/database"
	"github.com/GolangSriLanka/go-puso/internal/maths"
)

type Puso struct {
	gorm.Model `swaggerignore:"true"`
	Name       string  `json:"name,omitempty"`
	Color      string  `json:"color,omitempty"`
	Weight     float32 `json:"weight,omitempty"`
	Owner      string  `json:"owner,omitempty"`
	Laziness   int     `json:"laziness,omitempty"`
}

func (obj *Puso) Save() error {
	obj.Laziness = maths.LazinessGen(obj.Name, obj.Color, obj.Weight)

	if err := database.Database().Create(&obj).Error; err != nil {
		log.Errorf("failed to create puso: %+v: %v", obj, err)
		return err
	}

	return nil
}

func (obj *Puso) GetList() ([]Puso, error) {
	list := make([]Puso, 0)

	if err := database.Database().Find(&list).Error; err != nil {
		log.Errorf("failed to find puso: %+v: %v", obj, err)
		return nil, err
	}

	return list, nil
}

func (obj *Puso) Get(id string) (*Puso, error) {
	t := &Puso{}

	if err := database.Database().Where("ID = ?", id).First(t).Error; err != nil {
		log.Errorf("failed to find puso: %+v: %v", obj, err)
		return t, err
	}

	return t, nil
}

func (obj *Puso) Delete(id string) error {
	if err := database.Database().Model(&obj).Where("ID = ?", id).Delete(obj).Error; err != nil {
		log.Errorf("failed to delete puso: %+v: %v", obj, err)
		return err
	}

	return nil
}

func (obj *Puso) Update(id string) error {
	if err := database.Database().Model(&obj).Where("ID = ?", id).Updates(obj).Error; err != nil {
		log.Errorf("failed to update puso: %+v: %v", obj, err)
		return err
	}

	return nil
}

func (obj *Puso) Migrate() error {
	if err := database.Database().Migrator().AutoMigrate(obj); err != nil {
		log.Errorf("failed to migrate puso: %+v: %v", obj, err)
		return err
	}

	return nil
}
