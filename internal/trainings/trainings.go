package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("input string must contain 3 parts: steps,type,duration")
	}

	stepStr := strings.TrimSpace(parts[0])
	stepStr = strings.TrimPrefix(stepStr, "+") //trim добавлены, потому что не проходили автотесты
	steps, err := strconv.Atoi(stepStr)
	if err != nil {
		return fmt.Errorf("failed to parse steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be a positive integer")
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {

	var (
		distance = spentenergy.Distance(t.Steps, t.Height)
		speed    = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
		calories float64
		err      error
	)

	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("calorie calculation error: %w", err)
	}

	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return info, nil
}
