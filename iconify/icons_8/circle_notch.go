package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleNotch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 4.18v2.022c4.56.93 8 4.97 8 9.798c0 5.514-4.486 10-10 10S6 21.514 6 16c0-4.83 3.44-8.87 8-9.798v-2.02C8.334 5.136 4 10.065 4 16c0 6.617 5.383 12 12 12s12-5.383 12-12c0-5.934-4.334-10.863-10-11.82z"/>`),
		g.Group(children),
	)
}