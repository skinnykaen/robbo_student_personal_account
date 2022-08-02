package models

import (
	"gorm.io/gorm"
	"strconv"
)

type StudentCore struct {
	UserCore
}

type StudentHTTP struct {
	UserHttp `json:"userHttp"`
}

type StudentDB struct {
	gorm.Model
	UserDB
}

func (em *StudentDB) ToCore() *StudentCore {
	return &StudentCore{
		UserCore: UserCore{
			Id:         strconv.FormatUint(uint64(em.ID), 10),
			Email:      em.Email,
			Password:   em.Password,
			Role:       Role(em.Role),
			Nickname:   em.Nickname,
			Firstname:  em.Firstname,
			Lastname:   em.Lastname,
			Middlename: em.Middlename,
			CreatedAt:  em.CreatedAt.String(),
		},
	}
}

func (em *StudentDB) FromCore(student *StudentCore) {
	id, _ := strconv.ParseUint(student.Id, 10, 64)
	em.ID = uint(id)
	em.Email = student.Email
	em.Password = student.Password
	em.Role = uint(student.Role)
	em.Nickname = student.Nickname
	em.Firstname = student.Firstname
	em.Lastname = student.Lastname
	em.Middlename = student.Middlename
}

func (ht *StudentHTTP) ToCore() *StudentCore {
	return &StudentCore{
		UserCore: UserCore{
			Id:         ht.Id,
			Email:      ht.Email,
			Password:   ht.Password,
			Role:       Role(ht.Role),
			Nickname:   ht.Nickname,
			Firstname:  ht.Firstname,
			Lastname:   ht.Lastname,
			Middlename: ht.Middlename,
		},
	}
}

func (ht *StudentHTTP) FromCore(student *StudentCore) {
	ht.Id = student.Id
	ht.CreatedAt = student.CreatedAt
	ht.Email = student.Email
	ht.Password = student.Password
	ht.Role = uint(student.Role)
	ht.Nickname = student.Nickname
	ht.Firstname = student.Firstname
	ht.Lastname = student.Lastname
	ht.Middlename = student.Middlename
}
