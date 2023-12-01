package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Alert struct {
	date    time.Time
	content string
	userID  string
	period  time.Duration
	alertID string
}

type AlertsDB struct {
	f      *os.File
	alerts []Alert
}

func NewAlertDB(path string) (*AlertsDB, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	a := &AlertsDB{
		f:      f,
		alerts: nil,
	}

	return a, nil
}

func (a *AlertsDB) AddAlert(alert Alert) error {

	a.alerts = append(a.alerts, alert)

	err := a.save()
	if err != nil {
		return err
	}

	return nil
}

func (a *AlertsDB) save() error {
	_, err := a.f.Seek(0, 2)
	if err != nil {
		return fmt.Errorf("can't reset file offset: %w", err)
	}

	err = a.f.Truncate(0)
	if err != nil {
		return fmt.Errorf("can't truncate the file: %w", err)
	}

	for i := range a.alerts {
		_, err := a.f.Write([]byte(fmt.Sprintf("%s,%s,%s,%s,%s\n",
			a.alerts[i].alertID,
			a.alerts[i].date,
			a.alerts[i].userID,
			a.alerts[i].content,
			a.alerts[i].period),
		))
		if err != nil {
			return fmt.Errorf("can't save string: %w", err)
		}
	}
	return a.f.Sync()
}

func (a *AlertsDB) load() error {
	reader := csv.NewReader(a.f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		date, err := time.Parse(time.RFC3339, record[0])
		if err != nil {
			return err
		}

		period, err := time.ParseDuration(record[4])
		if err != nil {
			return err
		}

		alert := Alert{
			date:    date,
			content: record[1],
			userID:  record[2],
			period:  period,
			alertID: record[5],
		}

		a.alerts = append(a.alerts, alert)
	}

	return nil
}

func (a *AlertsDB) GetAlerts(userIDToFind string) ([]Alert, error) {
	res := make([]Alert, 0)

	for i := range a.alerts {
		if a.alerts[i].userID == userIDToFind {
			res = append(res, a.alerts[i])
		}
	}

	return res, nil
}

func (a *AlertsDB) UpdateAlert(alert Alert) error {
	for i := range a.alerts {
		if a.alerts[i].alertID == alert.alertID {
			a.alerts[i] = alert
		}
	}
	return a.save()
}
