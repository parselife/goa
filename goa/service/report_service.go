package service

import "goa/goa/repo"

type ReportService interface {

}

type reportService struct {
	weeklyRepo  repo.WeeklyReportRepository
	monthlyRepo repo.MonthlyReportRepository
}

func NewReportService(r repo.WeeklyReportRepository, r1 repo.MonthlyReportRepository) ReportService {
	return &reportService{r, r1}
}
