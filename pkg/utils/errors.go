package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func RenderOr500(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	err := render.Render(w, r, v)
	if err != nil {
		_ = render.Render(w, r, NewErrorResponse(errors.Wrap(err, "problem rendering response"), 500))
	}
}

func RenderListOr500(w http.ResponseWriter, r *http.Request, vs []render.Renderer) {
	err := render.RenderList(w, r, vs)
	if err != nil {
		_ = render.Render(w, r, NewErrorResponse(errors.Wrap(err, "problem rendering response"), 500))
	}
}

type ErrorResponse struct {
	Error      string   `json:"error"`
	Cause      string   `json:"cause"`
	StackTrace []string `json:"stacktrace"`
	Code       int      `json:"code"`
	Text       string   `json:"text"`
}

func (e *ErrorResponse) Render(h http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Code)
	return nil
}

func NewErrorResponse(err error, code int) *ErrorResponse {
	if err != nil {
		logrus.Error(err)
	}
	return &ErrorResponse{
		Error:      err.Error(),
		Cause:      errors.Cause(err).Error(),
		StackTrace: splitTrace(err),
		Code:       code,
		Text:       http.StatusText(code),
	}
}

func splitTrace(err error) []string {
	text := fmt.Sprintf("%+v", err)
	tmpLines := strings.Split(text, "\n")
	lines := make([]string, 0)
	for _, line := range tmpLines {
		lines = append(lines, strings.Replace(line, "\t", "        ", 1))
	}
	return lines
}
