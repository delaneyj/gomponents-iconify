package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20H8a4.92 4.92 0 0 0 1-3v-3h6a1 1 0 0 0 0-2H9V8.89a4.89 4.89 0 0 1 9.13-2.44a1 1 0 0 0 1.37.36a1 1 0 0 0 .37-1.36A6.9 6.9 0 0 0 7 8.89V12H4a1 1 0 0 0 0 2h3v3a3 3 0 0 1-3 3a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}