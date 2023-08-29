package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalanceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m16 22l-6-10l-6 10"/><path d="M10 28a6 6 0 0 0 6-6H4a6 6 0 0 0 6 6Z" clip-rule="evenodd"/><path d="m44 22l-6-10l-6 10"/><path d="M38 28a6 6 0 0 0 6-6H32a6 6 0 0 0 6 6Z" clip-rule="evenodd"/><path d="M24 6v36M10 12h28m-28 0h28m0 30H10"/></g>`),
		g.Group(children),
	)
}