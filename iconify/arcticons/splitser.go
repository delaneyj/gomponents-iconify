package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Splitser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.75 34.6a3.66 3.66 0 0 1-2.58-1.06l-7.46-7.47a3.65 3.65 0 0 1 5.16-5.16l4.88 4.88l11.52-11.51a3.65 3.65 0 1 1 5.16 5.16l-14.09 14.1a3.66 3.66 0 0 1-2.59 1.06Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.57 26.26a3.65 3.65 0 0 1 5.16-5.17l7.47 7.47a3.65 3.65 0 1 1-5.2 5.16Z"/>`),
		g.Group(children),
	)
}