package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 11.32c0 2.584-2.144 4.68-4.787 4.68H3.617C1.619 16 0 14.416 0 12.463c0-1.951 1.619-3.535 3.617-3.535c.146 0 .288.012.429.027a5.076 5.076 0 0 1-.057-.756C3.989 5.328 6.37 3 9.309 3c2.407 0 4.439 1.562 5.096 3.707a5 5 0 0 1 .809-.066C17.856 6.641 20 8.734 20 11.32z"/>`),
		g.Group(children),
	)
}