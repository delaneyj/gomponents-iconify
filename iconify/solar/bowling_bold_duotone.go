package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlingBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z" opacity=".5"/><path d="M12 10.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM8 8a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4-2.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/></g>`),
		g.Group(children),
	)
}