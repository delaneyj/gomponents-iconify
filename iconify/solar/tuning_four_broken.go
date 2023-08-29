package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningFourBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm2-8a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-4 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-linecap="round" d="M17.5 20H19m-7 0h2.75M6.5 4H5m7 0H9.25M19 12h-5m5-8h-3M5 20h2.667M10 12H7.5m-2 0H5"/></g>`),
		g.Group(children),
	)
}