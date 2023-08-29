package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m14.5 14.5l-4-4M4 6.5h5m-2.5 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		g.Group(children),
	)
}