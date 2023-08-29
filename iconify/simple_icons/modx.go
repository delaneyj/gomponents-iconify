package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.145 9.49l4.47-7.19H12.118l-1.24 2.023zM2.92 0v11.497l2.48 1.55l13.435-3.1zm18.16 24V12.503l-1.984-1.263l-5.168 8.267zM5.165 14.053l-4.78 7.648h11.497L18.6 10.953Z"/>`),
		g.Group(children),
	)
}