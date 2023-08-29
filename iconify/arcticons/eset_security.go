package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EsetSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.595 0 14.7-12.374 15.094-20.26l-1.419-2.026l1.52-2.026V7.22L26.026 4.5L24 5.21l-2.026-.71L8.805 8.106V19.25l1.52 1.965l-1.418 2.026C9.385 32.54 22.417 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.974 4.5v8.104l-9.117 6.584H8.805M26.026 4.5v8.104l9.117 6.584h4.052m-.102 4.052l-4.76.06L24 15.44l-10.332 7.8H8.907M24 15.44V43.5"/>`),
		g.Group(children),
	)
}