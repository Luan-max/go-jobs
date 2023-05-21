package handler

import (
	"fmt"
)

func validatePropsIsRequire(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateJobRequest struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Salary      string `json:"salary"`
	Benefits    string `json:"benefits"`
}

func (r *CreateJobRequest) Validate() error {
	if r.Title == "" {
		return validatePropsIsRequire("title", "string")
	}
	if r.Company == "" {
		return validatePropsIsRequire("company", "string")
	}
	if r.Salary == "" {
		return validatePropsIsRequire("salary", "string")
	}
	if r.Description == "" {
		return validatePropsIsRequire("description", "string")
	}
	if r.Link == "" {
		return validatePropsIsRequire("link", "string")
	}
	if r.Benefits == "" {
		return validatePropsIsRequire("benefits", "string")
	}
	if r.Remote == nil {
		return validatePropsIsRequire("remote", "bool")
	}
	return nil
}
