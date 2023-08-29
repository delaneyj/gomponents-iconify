package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeAl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#2ea98e" d="M11.616 7.986A1.559 1.559 0 0 0 10.16 7h-.06a1.558 1.558 0 0 0-1.456.986L2 25h3.806l1.015-2.834h6.621L14.457 25h3.8Zm-3.672 10.97l2.188-6.111l2.188 6.116Zm15.885 2.715V7.129H20.3v15.618A2.346 2.346 0 0 0 22.57 25H30v-3.328Z"/>`),
		g.Group(children),
	)
}