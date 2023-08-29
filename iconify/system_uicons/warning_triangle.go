package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(1 1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5.5l9 16H.5zm0 10v-5"/><circle cx="9.5" cy="13.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}