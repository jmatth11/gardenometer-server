package actions

import (
	"errors"
	"gardenometer/models"
)

type Queue struct {
  data []models.Action
}

func (q *Queue) Push(s models.Action) {
  q.data = append(q.data, s)
}

func (q *Queue) Pop() (models.Action, error) {
  if len(q.data) == 0 {
    return models.Action{}, errors.New("empty queue")
  }
  a := q.data[0]
  q.data = q.data[1:]
  return a, nil
}

func (q *Queue) Len() int {
  return len(q.data)
}
