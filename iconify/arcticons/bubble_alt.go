package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.152 5.5L42.5 34.847L34.846 42.5L5.5 13.153z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.2 16.546a3.907 3.907 0 0 0-1.043.745c-1.771 1.771-1.517 4.898.569 6.983s5.212 2.34 6.983.569a3.905 3.905 0 0 0 .745-1.042m-16.412-1.105l7.653-7.653m2.609 17.915l7.653-7.653"/>`),
		g.Group(children),
	)
}