package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatsDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 14h14v2H0V0h2zm2.5-1a1.5 1.5 0 1 1 .131-2.994l1.612-2.687a1.5 1.5 0 1 1 2.514 0l1.612 2.687a1.42 1.42 0 0 1 .23-.002l2.662-4.658a1.5 1.5 0 1 1 1.14.651l-2.662 4.658a1.5 1.5 0 1 1-2.496.026L7.631 7.994a1.42 1.42 0 0 1-.262 0l-1.612 2.687A1.5 1.5 0 0 1 4.5 13z"/>`),
		g.Group(children),
	)
}