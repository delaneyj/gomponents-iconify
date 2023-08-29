package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.834 14.693a7.062 7.062 0 0 0 1.193-7.334l-3.646 4.25l-3.593-.698l-1.19-3.462l3.636-4.242a7.063 7.063 0 0 0-8.106 9.903L5.625 24.04a2.79 2.79 0 0 0 4.209 3.661l9.493-10.917a7.072 7.072 0 0 0 7.508-2.09z"/>`),
		g.Group(children),
	)
}