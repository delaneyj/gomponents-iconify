package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSensorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5h2a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-2M7 19H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h2m1 7a4 4 0 1 0 8 0a4 4 0 1 0-8 0"/>`),
		g.Group(children),
	)
}