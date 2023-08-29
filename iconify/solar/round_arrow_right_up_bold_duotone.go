package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundArrowRightUpBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z" opacity=".5"/><path d="M10.5 8.25a.75.75 0 0 0 0 1.5h2.69l-4.72 4.72a.75.75 0 1 0 1.06 1.06l4.72-4.72v2.69a.75.75 0 0 0 1.5 0V9a.75.75 0 0 0-.75-.75h-4.5Z"/></g>`),
		g.Group(children),
	)
}