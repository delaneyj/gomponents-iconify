package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinwheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M21 4v17H11L21 4Zm6 40V27h10L27 44Zm0-33l17 10H27V11Zm-6 26L4 27h17v10Z"/>`),
		g.Group(children),
	)
}