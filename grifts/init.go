package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/wedneyyuri/buffalosample/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
