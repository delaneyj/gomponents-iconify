package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5A21.5 21.5 0 1 1 45.5 24A21.51 21.51 0 0 1 24 45.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.75 14.4A7.09 7.09 0 0 1 9 13.09c8.44 0 7.91 22 18.09 22c8.82 0 8.31-18.09 16.63-19.77"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.23 13.51c6.52 3 6.84 21.58 16.18 21.58S29.66 15.2 39.33 15.2a6.87 6.87 0 0 1 5.17 2.3"/>`),
		g.Group(children),
	)
}