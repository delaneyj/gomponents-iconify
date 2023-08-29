package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitterAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.74 16.55v1c0 10.07-7.64 21.61-21.62 21.61A21.14 21.14 0 0 1 5.5 35.71a12.22 12.22 0 0 0 1.81.11a15.25 15.25 0 0 0 9.44-3.24a7.56 7.56 0 0 1-7.1-5.29a6.9 6.9 0 0 0 1.44.15a7.53 7.53 0 0 0 2-.27A7.57 7.57 0 0 1 7 19.72v-.1a7.42 7.42 0 0 0 3.44.94A7.54 7.54 0 0 1 8.05 10.5a21.58 21.58 0 0 0 15.68 7.94a6.38 6.38 0 0 1-.21-1.74a7.55 7.55 0 0 1 13.17-5.31a15.59 15.59 0 0 0 4.83-1.85a7.65 7.65 0 0 1-3.39 4.27a15.87 15.87 0 0 0 4.37-1.26a15.56 15.56 0 0 1-3.76 4Z"/>`),
		g.Group(children),
	)
}