package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkypeMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1591 823q42 0 78 16t64 43t43 63t16 79q0 42-16 78t-43 64t-63 43t-79 16H395q-42 0-78-16t-64-43t-42-63t-16-79q0-42 15-78t43-64t63-43t79-16h1196z"/>`),
		g.Group(children),
	)
}