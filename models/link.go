package models

import "gorm.io/gorm"

type Bulb struct {
	gorm.Model
	UID   string `json:"uid"`
	Type  uint8  `json:"type"`
	User  uint64 `json:"user"`
	Links []Link `json:"links"`
}

type Link struct {
	gorm.Model
	Name   string `json:"name"`
	Link   string `json:"link"`
	BulbID uint   `json:"_"`
}

type BulbDTO struct {
	Type  uint8     `json:"type"`
	Links []LinkDTO `json:"links"`
}

func (bDTO *BulbDTO) LinksDTOToLinks() []Link {
	var links []Link
	for _, linkDTO := range bDTO.Links {
		link := Link{
			Name: linkDTO.Name,
			Link: linkDTO.Link,
		}
		links = append(links, link)
	}
	return links
}

type LinkDTO struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
