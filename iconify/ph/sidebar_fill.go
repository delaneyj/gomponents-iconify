package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SidebarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16ZM64 152H48a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm0-32H48a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm0-32H48a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm152 112H88V56h128v144Z"/>`),
		g.Group(children),
	)
}