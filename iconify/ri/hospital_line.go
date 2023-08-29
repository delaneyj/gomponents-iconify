package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20v-6h8v6h3V4H5v16h3Zm2 0h4v-4h-4v4Zm11 0h2v2H1v-2h2V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v17ZM11 8V6h2v2h2v2h-2v2h-2v-2H9V8h2Z"/>`),
		g.Group(children),
	)
}