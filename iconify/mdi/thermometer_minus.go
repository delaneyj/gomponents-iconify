package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11h-8v2h8v-2m-11 2V5c0-1.66-1.34-3-3-3S5 3.34 5 5v8c-2.21 1.66-2.66 4.79-1 7s4.79 2.66 7 1s2.66-4.79 1-7a4.74 4.74 0 0 0-1-1M8 4c.55 0 1 .45 1 1v3H7V5c0-.55.45-1 1-1Z"/>`),
		g.Group(children),
	)
}