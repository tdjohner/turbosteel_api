package models

import (
	"time"

	"gorm.io/gorm"
)

type Dog struct {
	gorm.Model
	Name      string
	BirthDate time.Time
	Facility  uint
}

//What are all the things I want to know about the organization
type Inistitution struct {
	gorm.Model
	Name string
}

//The physical building. Always owned by an instituion
type Facility struct {
	gorm.Model
	Inistitution Inistitution
	Address      string
	PostalCode   PostalCode
	City         City
	Province     Province
	Country      Country
}

type ImageFile struct {
	gorm.Model
	Dog      Dog
	FilePath string
}

type MedicalRecord struct {
	gorm.Model
	Dog      Dog
	Filepath string
}

type PostalCode struct {
	gorm.Model
	PostalCode string
	Country    Country
}

type City struct {
	gorm.Model
	Name     string
	Province Province
}

type Province struct {
	gorm.Model
	Name    string
	Country Country
}

type Country struct {
	gorm.Model
	Name string
}
