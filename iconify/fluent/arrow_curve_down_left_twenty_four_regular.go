package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveDownLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.401 3.378a.75.75 0 0 0-1.023-.28c-2.269 1.297-3.391 2.954-3.921 4.71c-.468 1.553-.463 3.166-.458 4.543l.001.4v5.688l-3.72-3.72a.75.75 0 1 0-1.06 1.061l5 5a.75.75 0 0 0 1.06 0l5-5a.75.75 0 0 0-1.06-1.06l-3.72 3.72v-6.07c-.004-1.411-.007-2.802.393-4.128c.42-1.394 1.298-2.737 3.23-3.84a.75.75 0 0 0 .278-1.024Z"/>`),
		g.Group(children),
	)
}