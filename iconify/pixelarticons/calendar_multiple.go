package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2h2v2h4v14H5V4h4V2h2v2h6V2zm-6 4H7v2h14V6H11zm-4 4v6h14v-6H7zM3 20h16v2H1V8h2v12z"/>`),
		g.Group(children),
	)
}