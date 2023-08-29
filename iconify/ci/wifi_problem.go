package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiProblem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.343 14.59a4.998 4.998 0 0 1 7.29-.025m-9.484-3.021A8 8 0 0 1 11.96 9m-8.735-.184A12 12 0 0 1 11.959 5M16 9l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m-8 10a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}