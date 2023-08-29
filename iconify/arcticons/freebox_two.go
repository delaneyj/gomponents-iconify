package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeboxTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.73 42.11h-9.45c-9.45 0-14.74-10.4-10-18.08l7.15-11.68c2.18-3.56 6-6.5 10.27-6.46h0c3.33 0 6.75 2 9.4 6.47L41 24a12.82 12.82 0 0 1 .34 11.41M24.66 5.89"/>`),
		g.Group(children),
	)
}