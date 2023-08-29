package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20h2v2H1v-2h2V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v17ZM11 8H9v2h2v2h2v-2h2V8h-2V6h-2v2Zm3 12h2v-6H8v6h2v-4h4v4Z"/>`),
		g.Group(children),
	)
}