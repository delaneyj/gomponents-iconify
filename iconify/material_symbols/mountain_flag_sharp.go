package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainFlagSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 12.875l-1.925.975l-3.05-2.025L8.4 9H11V2h7l-1 2l1 2h-5v3h2.5l1.425 2.85l-3 2L12 12.875ZM2 22l4.125-8.375l3.8 2.525L12 15.125l2.075 1.025l3.75-2.475L22 22H2Z"/>`),
		g.Group(children),
	)
}