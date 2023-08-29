package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSolidMirroredTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M294 0q-61 0-114 21T86 81t-63 91T0 286q0 47 10 84t28 69t44 62t58 61l424-423q-31-31-59-57t-60-44t-67-28t-84-10zm420 290L292 712l1003 1001l753 333l-348-759l-986-997z"/>`),
		g.Group(children),
	)
}