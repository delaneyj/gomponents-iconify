package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TownHallEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0L1 2v1h9V2L5.5 0zM2 4v4L1 9v1h9V9L9 8V4H2zm1 1h1v3H3V5zm2 0h1v3H5V5zm2 0h1v3H7V5z" fill="currentColor"/>`),
		g.Group(children),
	)
}