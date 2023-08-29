package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMinimalisticTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m14.162 18.488l-.72.72a6.117 6.117 0 0 1-8.65-8.65l.72-.72" opacity=".5"/><path d="m9.837 14.162l4.325-4.325"/><path d="m9.837 5.512l.721-.72a6.117 6.117 0 1 1 8.65 8.65l-.72.72" opacity=".5"/></g>`),
		g.Group(children),
	)
}