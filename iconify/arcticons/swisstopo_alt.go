package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwisstopoAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.939 21.908v-7.155m-3.585 3.585h7.17m-3.51 8.34c7.065-2.715 6.705-8.13 6.66-11.4v-2.46a25.214 25.214 0 0 0-6.66-1.14a25.214 25.214 0 0 0-6.66 1.14v2.46c-.12 3.27-.48 8.685 6.66 11.4Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a13.09 13.09 0 0 0-13.09 13.09h0c0 10.25 10 22.61 12.61 25.63a.8.8 0 0 0 1.21 0c2.55-3 12.36-15.38 12.36-25.63A13.09 13.09 0 0 0 24 4.5Z"/>`),
		g.Group(children),
	)
}