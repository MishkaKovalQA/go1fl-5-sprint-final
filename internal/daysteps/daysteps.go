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
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("invalid input format: expected 2 parts, got %d: %w", len(slice), err)
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be greater than zero: %w", err)
	}

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be greater than zero: %w", err)
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error calculating calories: %w", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return result, nil
}
