package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeFauna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#323fcb" d="m4.682 2l14.035 11.191l2.428-3.122l6.173-1.248l-3.118 5.133v6.589L7.341 30l4.439-7.237L16 20.058l-4.22.971l-3.953-6.451L15.6 16.59l-8.67-4.232Z"/>`),
		g.Group(children),
	)
}