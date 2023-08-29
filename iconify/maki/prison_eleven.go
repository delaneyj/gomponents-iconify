package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrisonEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M2 1v9h7V1H2zm1 1h1v3H3V2zm2 0h1v3H5V2zm2 0h1v2H7V2zm.5 3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1zM3 6h1v3H3V6zm2 0h1v3H5V6zm2 1h1v2H7V7z" fill="currentColor"/>`),
		g.Group(children),
	)
}