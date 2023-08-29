package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noodles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24c0 11.046 6.667 20 20 20s20-8.954 20-20H4Z" clip-rule="evenodd"/><path d="M16 24V8m8 16V6m8 18V4M8 24V10m-4 3l40-9"/></g>`),
		g.Group(children),
	)
}