package routes

import "gardenometer/models"

func ErrorToast(msg string) models.Toast {
  return models.Toast{
    Message: msg,
    ClassName: "notification is-danger",
  }
}

func SuccessToast(msg string) models.Toast {
  return models.Toast{
    Message: msg,
    ClassName: "notification is-primary",
  }
}
