package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 16H4v26h40V16Zm0 0V6H4v10h40ZM26 6l4 10M18 6l4 10M10 6l4 10M34 6l4 10m-26 8h24m-24 8h12"/>`),
		g.Group(children),
	)
}