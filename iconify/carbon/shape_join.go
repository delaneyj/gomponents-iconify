package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeJoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 10h-6V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v16a2.002 2.002 0 0 0 2 2h6v6a2.002 2.002 0 0 0 2 2h16a2.002 2.002 0 0 0 2-2V12a2.002 2.002 0 0 0-2-2ZM4 20V4h6v16Zm18 8V12h6v16Z"/>`),
		g.Group(children),
	)
}