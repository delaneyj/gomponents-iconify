package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentificationBadgeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16ZM96 48h64a8 8 0 0 1 0 16H96a8 8 0 0 1 0-16Zm84.81 150.4a8 8 0 0 1-11.21-1.6a52 52 0 0 0-83.2 0a8 8 0 1 1-12.8-9.6a67.88 67.88 0 0 1 27.4-21.69a40 40 0 1 1 53.94 0a67.88 67.88 0 0 1 27.46 21.69a8 8 0 0 1-1.59 11.2ZM152 136a24 24 0 1 1-24-24a24 24 0 0 1 24 24Z"/>`),
		g.Group(children),
	)
}