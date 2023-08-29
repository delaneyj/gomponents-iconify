package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarWeekend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h18V4h-4V2zM7 6h12v2H5V6h2zM5 20V10h14v10H5zm12-8h-2v6h2v-6z"/>`),
		g.Group(children),
	)
}