package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("invalid input format: expected 2 values")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("invalid number of steps: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps should be a positive number")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("invalid duration: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration should be a positive value")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("steps should be a positive number")
	}

	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration should be a positive value")
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories)

	return result, nil
}
