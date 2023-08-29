package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 15a6 6 0 0 1-6-6V1h12v8a6 6 0 0 1-6 6ZM6 3H1v4c0 2.509 1.791 4 4 4h1V3Zm12 8h1c2.209 0 4-1.491 4-4V3h-5v8ZM5 23h14v-4H5v4Zm11-4a4 4 0 1 0-8 0"/>`),
		g.Group(children),
	)
}