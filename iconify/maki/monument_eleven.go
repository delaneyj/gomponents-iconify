package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0L4 2v4.5h3V2L5.5 0zM3 7L2 8v3h7V8L8 7H3z" fill="currentColor"/>`),
		g.Group(children),
	)
}