package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeCTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#005f91" d="M23 19.418a6.971 6.971 0 1 1-.05-6.918l6.093-3.509a14 14 0 1 0 .036 13.95Z"/>`),
		g.Group(children),
	)
}