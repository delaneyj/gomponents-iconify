package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.172 9.172l-3.536 3.535a4 4 0 0 0 0 5.657v0a4 4 0 0 0 5.657 0l3.535-3.536M9.172 9.172l3.535-3.536a4 4 0 0 1 5.657 0v0a4 4 0 0 1 0 5.657l-3.536 3.535M9.172 9.172l5.656 5.656"/>`),
		g.Group(children),
	)
}