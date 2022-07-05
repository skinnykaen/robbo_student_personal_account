package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CourseDelegateImpl struct {
	courses.UseCase
}

type CourseDelegateModule struct {
	fx.Out
	courses.Delegate
}

func SetupCourseDelegate(usecase courses.UseCase) CourseDelegateModule {
	return CourseDelegateModule{
		Delegate: &CourseDelegateImpl{usecase},
	}
}

func (p *CourseDelegateImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, err error) {
	return p.UseCase.CreateCourse(course, courseId)
}

func (p *CourseDelegateImpl) DeleteCourse(course *models.CourseHTTP) (err error) {
	courseCore := course.ToCore()
	return p.UseCase.DeleteCourse(courseCore)
}

func (p *CourseDelegateImpl) GetCoursesByUser(username string) (body string, err error) {
	return p.UseCase.GetCoursesByUser(username)
}

func (p *CourseDelegateImpl) GetCourseContent(courseId string) (body string, err error) {
	return p.UseCase.GetCourseContent(courseId)
}

func (p *CourseDelegateImpl) GetAllPublicCourses(pageNumber int) (body string, err error) {
	return p.UseCase.GetAllPublicCourses(pageNumber)
}

func (p *CourseDelegateImpl) UpdateCourse(course *models.CourseHTTP) (err error) {
	courseCore := course.ToCore()
	return p.UseCase.UpdateCourse(courseCore)
}
