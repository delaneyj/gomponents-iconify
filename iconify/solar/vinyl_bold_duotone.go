package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z" opacity=".5"/><path fill-rule="evenodd" d="M13.818 3.395a.75.75 0 0 1 .669-.11a9.541 9.541 0 0 1 6.228 6.228a.75.75 0 1 1-1.43.45a8.045 8.045 0 0 0-4.273-4.87v7.276a3.381 3.381 0 1 1-1.5-2.81V4a.75.75 0 0 1 .306-.605Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}