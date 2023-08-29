package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypePddlPlan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="green" d="M4.178 4.956h17.37v6.45H4.178zm7.145 8.538h17.37v6.45h-17.37zm10.334 8.539h7.036v6.45h-7.036z"/>`),
		g.Group(children),
	)
}