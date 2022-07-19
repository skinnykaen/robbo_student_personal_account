package users

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	GetUsersByRole(role *models.Role) (user *[]models.UserHttp, err error)
	PostEnrollment(postEnrollmentHTTP *models.PostEnrollmentHTTP) (err error)
	PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error)
	UpdateUser(userHTTP *models.UserHttp) (err error)
	DeleteUser(userId int) (err error)
	RegistrationUser(userForm *edxApi.RegistrationForm) (err error)
	LoginUser(email, password string) (err error)
	GetUserById(userId int) (user models.UserHttp, err error)
	CreateUser(user *models.UserHttp) (id string, err error)
}