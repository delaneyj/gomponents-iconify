package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 2h-2v2H9V2H7v2H3v18h4v-2H5V10h14v10h-2v2h4V4h-4V2zM7 6h12v2H5V6h2zm6 6h-2v6H9v-2H7v2h2v2h2v2h2v-2h2v-2h2v-2h-2v2h-2v-6z"/>`),
		g.Group(children),
	)
}