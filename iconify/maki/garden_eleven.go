package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GardenEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 6a4 4 0 0 1-4.25 4A4 4 0 0 1 1.5 6a4 4 0 0 1 3.9 2.26V5h-2a.89.89 0 0 1-.9-.88V1.84a.35.35 0 0 1 .64-.21L4.28 3L5.45.67a.35.35 0 0 1 .6 0L7.22 3l1.13-1.38a.35.35 0 0 1 .65.21v2.28a.89.89 0 0 1-.89.89H6.1v3.26A4 4 0 0 1 10 6z" fill="currentColor"/>`),
		g.Group(children),
	)
}