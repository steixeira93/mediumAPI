package handler

import "fmt"

type CreateTextRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("Param: %s (type: %s) is required", name, typ)
}

func (r *CreateTextRequest) Validate() error {
	if r.Title == "" {
		return errParamsIsRequired("title", "string")
	}

	if r.Description == "" {
		return errParamsIsRequired("description", "string")
	}

	if r.Author == "" {
		return errParamsIsRequired("author", "string")
	}

	return nil
}
