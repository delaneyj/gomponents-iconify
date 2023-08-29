package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M21 13a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 13V9"/><path stroke-linecap="round" d="M10 2h4"/></g>`),
		g.Group(children),
	)
}