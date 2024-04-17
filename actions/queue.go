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

func (q *Queue) NextActionForName(name string) (models.Action, bool) {
  res := models.Action{}
  targetIndex := -1
  for i, e := range q.data {
    if e.Name == name {
      targetIndex = i
      break
    }
  }
  if targetIndex == -1 {
    return res, false
  }
  target := q.data[targetIndex]
  newOffset := targetIndex + 1
  if newOffset > len(q.data) {
    q.data = q.data[:targetIndex]
  } else {
    q.data = append(q.data[:targetIndex], q.data[targetIndex+1:]...)
  }
  return target, true
}
