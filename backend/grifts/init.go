package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/twatzl/StudySpaceFinder/backend/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
