//models

package models

import (
	"time"

	"gorm.io/gorm"
)

// What are all the things I want to know about the organization
type Institution struct {
	gorm.Model
	Name string `json:"institutionName"`
}

type ImageFile struct {
	gorm.Model
	Dog      Dog    `json:"dogID"`
	FilePath string `json:"imageFilePath"`
}

type RecordFiles struct {
	gorm.Model
	Dog      Dog    `json:"dogID"`
	Filepath string `json:"recordFilePath"`
}

type Country struct {
	gorm.Model
	Name string `json:"countryName"`
}

type PostalCode struct {
	gorm.Model
	PostalCode string `json:"postalCode"`
	CountryID  int
	Country    Country `json:"countryID" gorm:"foreignKey:CountryID;references:ID"`
}

type City struct {
	gorm.Model
	Name       string `json:"cityName"`
	ProvinceID int
	Province   Province `json:"provinceID" gorm:"foreignKey:ProvinceID;references:ID"`
}

type Province struct {
	gorm.Model
	Name      string `json:"provinceName"`
	CountryID int
	Country   Country `json:"countryID" gorm:"foreignKey:CountryID;references:ID"`
}

// The physical building. Always owned by an instituion
type Location struct {
	gorm.Model
	Name          string `json:"locationName"`
	Address       string `json:"locationAddress"`
	InstitutionID int
	PostalCodeID  int
	CityID        int
	ProvinceID    int
	Institution   Institution `json:"institutionID" gorm:"foreignKey:InstitutionID;references:ID"`
	PostalCode    PostalCode  `json:"postalCodeID" gorm:"foreignKey:PostalCodeID;references:ID"`
	City          City        `json:"cityID" gorm:"foreignKey:CityID;references:ID"`
	Province      Province    `json:"provinceID" gorm:"foreignKey:ProvinceID;references:ID"`
}

type Dog struct {
	gorm.Model
	Name       string    `json:"dogName"`
	BirthDate  time.Time `json:"dogBirthDate"`
	LocationID int
	Location   Location `json:"dogLocation" gorm:"foreignKey:LocationID;;references:ID"`
}
