package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPinFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.207 4.793A8.778 8.778 0 0 0 5.793 17.207l5.5 5.5a1 1 0 0 0 1.414 0l5.5-5.5a8.778 8.778 0 0 0 0-12.414ZM10.5 11A1.5 1.5 0 0 1 12 9.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5V11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}