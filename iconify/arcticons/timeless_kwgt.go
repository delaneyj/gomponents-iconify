package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimelessKwgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M33.43 43.327A21.413 21.413 0 0 1 24 45.5C12.126 45.5 2.5 35.874 2.5 24S12.126 2.5 24 2.5S45.5 12.126 45.5 24a21.4 21.4 0 0 1-2.174 9.431"/><circle cx="38.5" cy="38.5" r="7"/><path d="M36.654 34.492v8.016m3.692 0v-2.053L38.273 38.5l2.073-1.955v-2.053"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.123 29.127a7.25 7.25 0 1 0-1.976-6.588l-.294 2.922a7.25 7.25 0 1 1-1.976-6.588"/>`),
		g.Group(children),
	)
}