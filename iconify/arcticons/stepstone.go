package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stepstone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.58 34.32a.22.22 0 0 1-.26-.15A9 9 0 0 1 5 31.49c0-3.17 1.83-7.38 4-11.55c3.44-6.7 6.79-13.13 15.59-13.1c6.46.07 9.3 3.79 13 8.7c.91 1.24 1.77 2.3 2.55 3.19a.24.24 0 0 1 0 .33a106 106 0 0 1-16.44 9.32a169 169 0 0 1-18.12 5.94Zm10.05 5.31a15.17 15.17 0 0 1-7-1.42a.22.22 0 0 1-.1-.31a.29.29 0 0 1 .15-.12c4.38-1 12.18-2.86 18-4.58c6.3-1.83 13.32-6.83 16.54-9.29a.21.21 0 0 1 .31 0a.41.41 0 0 1 .05.09a9.87 9.87 0 0 1 .46 3.1C43.91 39 19 39.64 15.63 39.63Z"/>`),
		g.Group(children),
	)
}