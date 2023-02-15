package cohorts

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Delegate interface {
	CreateCohort(cohort *models.CohortHTTP, courseId string) (newCohort models.CohortHTTP, err error)
	AddStudent(username, courseId string, cohortId int) (err error)
}
