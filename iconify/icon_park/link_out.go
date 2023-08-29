package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 12V4H32V12"/><path d="M32 36V44H16V36"/><path d="M18 24L4 24"/><path d="M44 24L30 24"/><path d="M24 34V14"/><path d="M39 19L44 24L39 29"/><path d="M9 19L4 24L9 29"/></g>`),
		g.Group(children),
	)
}