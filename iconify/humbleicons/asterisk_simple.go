package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 4l.005 7.993m0 0l7.603-2.465m-7.603 2.465l4.697 6.48m-4.697-6.48l-4.707 6.48m4.707-6.48L4.392 9.528"/>`),
		g.Group(children),
	)
}