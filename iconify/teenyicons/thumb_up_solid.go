package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M9.312 2.995A2.034 2.034 0 0 0 5.776.992L3 5.354V12.5A2.5 2.5 0 0 0 5.5 15h5a2.5 2.5 0 0 0 2-1l2.5-3.333V7.5A2.5 2.5 0 0 0 12.5 5H8.309l1.003-2.005ZM0 5v10h1V5H0Z"/>`),
		g.Group(children),
	)
}