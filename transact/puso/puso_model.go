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

	if result := database.Database().Create(&obj); result.Error != nil {
		log.Errorf("failed to create puso: %+v: %v", obj, result.Error)
		return result.Error
	}

	return nil
}

func (obj *Puso) GetList() ([]Puso, error) {
	list := make([]Puso, 0)

	if result := database.Database().Find(&list); result.Error != nil {
		log.Errorf("failed to find puso: %+v: %v", obj, result.Error)
		return nil, result.Error
	}

	return list, nil
}

func (obj *Puso) Get(id string) (*Puso, error) {
	t := &Puso{}

	if result := database.Database().Where("ID = ?", id).First(t); result.Error != nil {
		log.Errorf("failed to find puso: %+v: %v", obj, result.Error)
		return t, result.Error
	}

	return t, nil
}

func (obj *Puso) Delete(id string) error {
	if result := database.Database().Model(&obj).Where("ID = ?", id).Delete(obj); result.Error != nil {
		log.Errorf("failed to delete puso: %+v: %v", obj, result.Error)
		return result.Error
	}

	return nil
}

func (obj *Puso) Update(id string) error {
	if result := database.Database().Model(&obj).Where("ID = ?", id).Updates(obj); result.Error != nil {
		log.Errorf("failed to update puso: %+v: %v", obj, result.Error)
		return result.Error
	}

	return nil
}

func (obj *Puso) Migrate() error {
	if err := database.Database().AutoMigrate(obj); err != nil {
		log.Errorf("failed to migrate puso: %+v: %v", obj, err)
		return err
	}
	return nil
}
