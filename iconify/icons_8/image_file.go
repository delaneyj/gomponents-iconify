package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8V5zm12 1.438L22.563 9H20V6.437zM21.094 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM14 15.594l-.72.687l-4 4l1.44 1.44L14 18.437l2.28 2.28l.72.688l.72-.687L19 19.437l2.28 2.28l1.44-1.437l-3-3l-.72-.686l-.72.687L17 18.563l-2.28-2.28l-.72-.688z"/>`),
		g.Group(children),
	)
}