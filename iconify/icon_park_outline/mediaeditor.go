package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mediaeditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M35 32.133V34a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-2.867L44 12"/><path d="M35 15.867V14a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v2.867L44 36"/></g>`),
		g.Group(children),
	)
}