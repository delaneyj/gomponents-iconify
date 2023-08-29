package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainFlagOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15.125Zm-3.3-2.2l1.375.925L12 12.875l1.925.975l1.325-.875l-1-1.975h-4.6l-.95 1.925ZM5.225 20H18.75l-2.6-5.225l-2.075 1.375L12 15.125L9.925 16.15L7.8 14.75L5.225 20ZM2 22L8.4 9H11V2h7l-1 2l1 2h-5v3h2.5L22 22H2Z"/>`),
		g.Group(children),
	)
}