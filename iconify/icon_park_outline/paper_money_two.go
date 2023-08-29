package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperMoneyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 13h40v24H4V13Z"/><path stroke-linecap="round" d="M4 21a8 8 0 0 0 8-8H4v8Zm0 8a8 8 0 0 1 8 8H4v-8Zm40 0v8h-8a8 8 0 0 1 8-8Zm0-8a8 8 0 0 1-8-8h8v8Z" clip-rule="evenodd"/><path d="M24 31c2.761 0 5-2.686 5-6s-2.239-6-5-6s-5 2.686-5 6s2.239 6 5 6Z"/></g>`),
		g.Group(children),
	)
}