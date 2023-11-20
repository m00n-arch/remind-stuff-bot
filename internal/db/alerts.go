package db

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Alert struct {
	date      time.Time
	content   string
	userID    string
	messageID int
	period    time.Duration
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

	return a, a.load()
}

func (a *AlertsDB) GetState(userIDToFind string) (string, error) {
	// for i := range a.alerts {
	// 	if a.alerts[i].userID == userIDToFind {
	// 		return a.alerts[i].state, nil
	// 	}
	// }
	return StartState, a.UpdateState(userIDToFind, StartState)
}

func (a *AlertsDB) UpdateState(userIDToFind string, newState string) error {
	// for i := range a.alerts {
	// 	if a.alerts[i].userID == userIDToFind {
	// 		a.alerts[i].state = newState
	// 		return a.save()
	// 	}
	// }
	// a.alerts = append(a.alerts, status{
	// 	userID: userIDToFind,
	// 	state:  newState,
	// })

	return a.save()
}

func (a *AlertsDB) save() error {
	_, err := a.f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("can't reset offset: %w", err)
	}

	err = a.f.Truncate(0)
	if err != nil {
		return fmt.Errorf("can't truncate the file: %w", err)
	}

	// for i := range a.alerts {
	// 	_, err := a.f.Write([]byte(fmt.Sprintf("%s,%s\n", a.alerts[i].userID, a.alerts[i].state)))
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return a.f.Sync()
}

func (a *AlertsDB) load() error {
	a.alerts = nil

	_, err := a.f.Seek(0, 0)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(a.f)
	if err != nil {
		return err
	}
	str := string(data)

	rows := strings.Split(str, "\n")

	for i := range rows {
		if i == len(rows)-1 {
			continue
		}

		// cells := strings.Split(rows[i], ",")
		// if len(cells) != 2 {
		// 	return fmt.Errorf("invalid cells count")
		// }
		// s := status{userID: cells[0], state: cells[1]}
		// a.alerts = append(a.alerts, s)
	}

	return nil
}

func accurate() {
	time.Parse("02.01.2006 15:04", "")
}
