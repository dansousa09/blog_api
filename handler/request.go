package handler

import "fmt"

func errFieldIsRequired(name, typ string) error {
	return fmt.Errorf("the %s field (%s) is required", name, typ)
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (r *CreatePostRequest) Validate() error {
	if r.Title == "" && r.Content == "" && r.Author == "" {
		return fmt.Errorf("the request body is unformatted or is empty")
	}

	if r.Title == "" {
		return errFieldIsRequired("title", "string")
	}

	if r.Content == "" {
		return errFieldIsRequired("content", "string")
	}

	if r.Author == "" {
		return errFieldIsRequired("author", "string")
	}

	return nil
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (r *UpdatePostRequest) Validate() error {
	if r.Title == "" || r.Content == "" || r.Author == "" {
		return fmt.Errorf("at least one of the Post fields are required")
	}

	return nil
}
