package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkRoundLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 9h-2a6 6 0 0 0 0 12h4a6 6 0 0 0 4.472-10" opacity=".5"/><path d="M10 15h2a6 6 0 0 0 0-12H8a6 6 0 0 0-4.472 10"/></g>`),
		g.Group(children),
	)
}