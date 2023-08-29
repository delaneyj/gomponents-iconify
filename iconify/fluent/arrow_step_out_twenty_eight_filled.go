package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowStepOutTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 16.5a.75.75 0 0 0 .75-.75V4.622l3.957 4.146a.75.75 0 0 0 1.085-1.036l-5.25-5.5a.75.75 0 0 0-1.085 0l-5.25 5.5a.75.75 0 0 0 1.086 1.036l3.957-4.146V15.75c0 .414.336.75.75.75Zm0 9.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}