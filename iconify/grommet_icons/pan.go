package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8.5 5.5L12 2l3.5 3.5M22 12H2m3.5-3.5L2 12l3.5 3.5m13 0L22 12l-3.5-3.5M12 22V2M8.5 18.5L12 22l3.5-3.5"/>`),
		g.Group(children),
	)
}