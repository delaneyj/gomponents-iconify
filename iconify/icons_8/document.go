package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V9.594l-.28-.313l-6-6l-.314-.28H6zm2 2h10v6h6v16H8V5zm12 1.438L22.563 9H20V6.437zM11 13v2h10v-2H11zm0 4v2h10v-2H11zm0 4v2h10v-2H11z"/>`),
		g.Group(children),
	)
}