package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14h6v6H2M16 8h-6v2h6M2 10h6V4H2m8 0v2h12V4M10 20h6v-2h-6m0-2h12v-2H10"/>`),
		g.Group(children),
	)
}