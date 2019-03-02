package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/twatzl/study_space_finder/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
