package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GardenCentreEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10.875 5.164l-.007-.008l-.029-.029l-.971-.97A.5.5 0 0 0 9 4.494v1.148L7 7.646V3a2 2 0 0 0-4 0a2 2 0 0 0 0 4v1a1 1 0 0 0 1 1h2a.984.984 0 0 0 .833-.48L9.36 5.99h1.061a.5.5 0 0 0 .453-.826zM1.5 5A1.5 1.5 0 0 1 3 3.5v3A1.5 1.5 0 0 1 1.5 5zm2-2a1.5 1.5 0 0 1 3 0h-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}