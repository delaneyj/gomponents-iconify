package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M43 12H3v30h40V12ZM11 6H3v6h8V6Zm32 0h-8v6h8V6ZM9 21h6m16 0h6m-25-3v6"/>`),
		g.Group(children),
	)
}