package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.697.04A.5.5 0 0 0 .04.697l6 14a.5.5 0 0 0 .934-.039l1.921-5.763l5.763-1.92a.5.5 0 0 0 .039-.935l-14-6Z"/>`),
		g.Group(children),
	)
}