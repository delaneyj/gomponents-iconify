package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m13.42 2.296l-8.047 6.94A.923.923 0 0 0 5 10c0 .315.137.58.41.797l8.01 6.907c.32.268.664.357 1.032.267c.3-.087.482-.29.548-.607v-14.7a.758.758 0 0 0-.526-.628c-.383-.091-.734-.005-1.054.26Z"/>`),
		g.Group(children),
	)
}