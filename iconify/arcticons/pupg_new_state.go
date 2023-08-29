package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PupgNewState(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.77 37.233C16.097 21.4 20.175 10.644 22.784 4.5c2.393 8.088 7.876 26.022 11.246 39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.658 22.265c7.94-5.142 16.317-9.102 26.684-11.07c-5.053 3.992-6.134 4.992-8.264 7.95"/>`),
		g.Group(children),
	)
}