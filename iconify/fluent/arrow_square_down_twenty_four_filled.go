package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 3A2.5 2.5 0 0 0 3 5.5v13A2.5 2.5 0 0 0 5.5 21h13a2.5 2.5 0 0 0 2.5-2.5v-13A2.5 2.5 0 0 0 18.5 3h-13Zm11.03 8.72a.75.75 0 0 1 .073.977l-.073.084l-4 4a.75.75 0 0 1-.977.073l-.084-.072l-4-4.002a.75.75 0 0 1 .976-1.133l.085.073l2.72 2.722V7.75a.75.75 0 0 1 .648-.743L12 7a.75.75 0 0 1 .743.648l.007.102v6.69l2.72-2.72a.75.75 0 0 1 .976-.072l.084.072Z"/>`),
		g.Group(children),
	)
}