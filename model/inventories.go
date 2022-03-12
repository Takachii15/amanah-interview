package model

import "github.com/Takachii/amanah-int/config"

type Inventories struct {
	ID       int `json:"id" form:"id" gorm:"primaryKey"`
	Nama     string `json:"nama" form:"nama"`
	Harga    int `json:"harga" form:"harga"`
	Stok    int `json:"stok" form:"stok"`
}

type Purchase struct {
  ID int `json:"id" form:"id" gorm:"foreignKey"`
  Purchase int `json:"purchase" form:"purchase"`
}

func (item *Inventories) CreateItem() error {
	if err := config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (item *Inventories) UpdateItem(id int) error {
	if err := config.DB.Model(&Inventories{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (item *Inventories) DeleteItem() error {
	if err := config.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}

func GetOneById(id int) (Inventories, error) {
	var item Inventories
	result := config.DB.Where("id = ?", id).First(&item)
	return item, result.Error
}

func GetAll() ([]Inventories, error) {
	var items []Inventories
	result := config.DB.Find(&items)

	return items, result.Error
}

