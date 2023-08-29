package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5V4H5v2H3v14h14V6h-2V4h-2v2H7V5zm-2 5V8h10v2H5zm0 2h10v6H5v-6zm16-3V8h-2v6h2V9zm0 6h-2v2h2v-2z"/>`),
		g.Group(children),
	)
}