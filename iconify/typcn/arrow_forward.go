package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5.499a.996.996 0 0 0-1 1v2.559c-4.5.498-8 4.309-8 8.941v1c2.245-3.423 5.25-3.92 8-3.989v2.489a.999.999 0 0 0 1.707.707L20 11.999l-6.293-6.208A.996.996 0 0 0 13 5.499z"/>`),
		g.Group(children),
	)
}