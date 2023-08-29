package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.793 37.1L9.434 26.965l9.562-12.494c1.837-2.4 4.111-2.442 6.502-1.147l13.068 7.075l-29.131 6.566"/>`),
		g.Group(children),
	)
}