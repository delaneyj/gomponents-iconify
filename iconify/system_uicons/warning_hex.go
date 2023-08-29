package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningHex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(-1 -1)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.517 3.5l4.983 5v6l-4.983 5H8.5l-5-5v-6l5-5zm-3.017 9v-5"/><circle cx="11.5" cy="15.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}