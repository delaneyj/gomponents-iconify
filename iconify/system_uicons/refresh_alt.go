package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 6.5c1.378-2.412 4.024-4 7-4a8 8 0 0 1 8 8m-1 4c-1.408 2.287-4.118 4-7 4a8 8 0 0 1-8-8"/><path d="M8.5 6.5h-5v-5m9 13h5v5"/></g>`),
		g.Group(children),
	)
}