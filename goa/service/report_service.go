package service

import "goa/goa/repo"

type ReportService interface {
}

type reportService struct {
	weeklyRepo  repo.WeeklyReportRepository
	monthlyRepo repo.MonthlyReportRepository
}
