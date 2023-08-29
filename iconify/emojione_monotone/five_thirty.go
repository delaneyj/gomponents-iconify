package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm2.564 47L34 46.896V58h-4V35.445c-1.19-.693-2-1.969-2-3.445c0-1.133.475-2.15 1.23-2.878l-.83-3.095L32.234 25l.269 1H34v2.554c1.19.692 2 1.968 2 3.446a3.986 3.986 0 0 1-1.139 2.789L38.4 47.975L34.564 49z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}