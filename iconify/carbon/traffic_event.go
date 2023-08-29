package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficEvent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 21h6v2h-6z"/><path fill="currentColor" d="m24.246 8l1.714 6H30v-2h-2.531l-1.3-4.549A2.008 2.008 0 0 0 24.246 6h-1.349l-.728-2.549A2.008 2.008 0 0 0 20.245 2H7.754a2.008 2.008 0 0 0-1.922 1.45L4.532 8H2v2h4.04l1.714-6h12.492l.571 2h-9.063a2.008 2.008 0 0 0-1.922 1.45L8.816 11H7.714a1.998 1.998 0 0 0-1.891 1.352L4.57 16H2v2h2v7a2.002 2.002 0 0 0 2 2v3h2v-3h12v3h2v-3a2.002 2.002 0 0 0 2-2v-7h2v-2h-2.571l-1.251-3.647A1.999 1.999 0 0 0 20.286 11h-9.389l.857-3ZM22 19v2h-2v2h2v2H6v-2h2v-2H6v-2Zm-.343-2H6.343l1.371-4h12.572Z"/>`),
		g.Group(children),
	)
}