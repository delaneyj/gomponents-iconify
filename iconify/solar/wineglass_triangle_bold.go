package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineglassTriangleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.3 3H4.7C3.2 3 2.439 4.794 3.485 5.864L6.235 9h11.53l2.751-3.136C21.562 4.794 20.8 3 19.3 3Zm-2.851 7.5H7.551l3.699 4.216v5.534H7.756a.75.75 0 1 0 0 1.5h8.488a.75.75 0 0 0 0-1.5H12.75v-5.533l3.699-4.217Z"/>`),
		g.Group(children),
	)
}