package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcludeSelection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="27" height="27" x="16" y="16" rx="2"/><rect width="27" height="27" x="5" y="5" rx="2"/><path d="m25 32l-9 9m25-25l-9 9m-16-2l-9 9M32 7l-9 9m20 8L24 43m0-38L5 24m38 10l-9 9M14 5l-9 9"/></g>`),
		g.Group(children),
	)
}