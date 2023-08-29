package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TickSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m11.39 5.312l-4.318 5.399L3.68 7.884l.64-.768l2.608 2.173l3.682-4.601l.78.624Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}