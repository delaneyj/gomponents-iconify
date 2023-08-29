package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.338 5.53c5.106 1.932 10.211 1.932 15.317 0A1 1 0 0 1 21 6.464v11c0 .692-.692 1.2-1.34.962c-5.107-1.932-10.214-1.932-15.321 0A.993.993 0 0 1 3 17.491V6.464a1 1 0 0 1 1.338-.935z"/>`),
		g.Group(children),
	)
}