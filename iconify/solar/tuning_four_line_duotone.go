package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningFourLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm2-8a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-4 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-linecap="round" d="M12 20h7m0-8h-5m5-8h-3m-4 0H5m5 8H5m0 8h2.667" opacity=".5"/></g>`),
		g.Group(children),
	)
}