package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleStroked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 12.3a4.8 4.8 0 1 1 0-9.6a4.8 4.8 0 0 1 0 9.6Zm0 1.7a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/>`),
		g.Group(children),
	)
}