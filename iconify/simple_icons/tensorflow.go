package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tensorflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.292 5.856L11.54 0v24l-4.095-2.378V7.603l-6.168 3.564l.015-5.31zm21.43 5.311l-.014-5.31L12.46 0v24l4.095-2.378V14.87l3.092 1.788l-.018-4.618l-3.074-1.756V7.603l6.168 3.564z"/>`),
		g.Group(children),
	)
}