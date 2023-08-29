package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roadblock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 14a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13ZM12 6v3H3V6h9Z"/>`),
		g.Group(children),
	)
}