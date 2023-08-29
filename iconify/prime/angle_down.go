package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14.5a.74.74 0 0 1-.53-.22L8 10.78a.75.75 0 0 1 1-1.06l3 3l3-3a.75.75 0 0 1 1 1.06l-3.5 3.5a.74.74 0 0 1-.5.22Z"/>`),
		g.Group(children),
	)
}