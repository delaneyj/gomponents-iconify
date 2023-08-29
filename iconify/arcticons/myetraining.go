package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myetraining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.077 26.715l-2.195 5.279l-32.418.123l2.197-5.285m2.261-5.433l2.19-5.282l32.352.03l-2.192 5.283"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.922 21.399l32.35.03l-2.195 5.286l-32.416.116l2.26-5.432Z"/>`),
		g.Group(children),
	)
}