package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CafeEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7 9.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1 0-1h4a.5.5 0 0 1 .5.5zM8 3H7V2H2v4a2.5 2.5 0 0 0 4.79 1H8a2 2 0 1 0 0-4zm0 3H7V4h1a1 1 0 1 1 0 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}