package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMorocco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#f42f4c"/><path fill="#4f682e" d="m23.3 44l3.3-9.9l-8.6-6.2h10.7L32 18l3.3 9.9H46L37.3 34l3.3 9.9l-8.6-6l-8.7 6.1m10.5-7.4l2.7 1.9l-1-3.1l-1.7 1.2m-5.3-1.2l-1 3.1l2.7-1.9l-1.7-1.2m.7-2.1l2.8 2l2.8-2l-1.1-3.2h-3.5l-1 3.2m-4.5-3.2l2.7 1.9l.6-1.9h-3.3M36.6 32l2.7-1.9H36l.6 1.9M31 27.9h2l-1-3.1l-1 3.1"/>`),
		g.Group(children),
	)
}