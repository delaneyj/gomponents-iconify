package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortFromBottomToTopBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17 20.75a.75.75 0 0 0 .75-.75V6.25l1.65 2.2a.75.75 0 1 0 1.2-.9l-3-4a.75.75 0 0 0-1.35.45v16c0 .414.336.75.75.75Z" clip-rule="evenodd"/><path d="M3.25 8c0 .414.336.75.75.75h9a.75.75 0 0 0 0-1.5H4a.75.75 0 0 0-.75.75Z"/><path d="M5.25 13c0 .414.336.75.75.75h7a.75.75 0 0 0 0-1.5H6a.75.75 0 0 0-.75.75Z" opacity=".7"/><path d="M7.25 18c0 .414.336.75.75.75h5a.75.75 0 0 0 0-1.5H8a.75.75 0 0 0-.75.75Z" opacity=".4"/></g>`),
		g.Group(children),
	)
}