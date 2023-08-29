package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12 3.75l7.274.205a2.584 2.584 0 0 1 2.494 2.29a37.436 37.436 0 0 1 0 8.51a2.584 2.584 0 0 1-2.494 2.29l-6.024.17v2.035H17a.75.75 0 1 1 0 1.5H7a.75.75 0 0 1 0-1.5h3.75v-2.035l-6.024-.17a2.584 2.584 0 0 1-2.494-2.29a37.434 37.434 0 0 1 0-8.51a2.584 2.584 0 0 1 2.494-2.29L12 3.75Zm-7 10a.75.75 0 1 0 0 1.5h14a.75.75 0 0 0 0-1.5H5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}