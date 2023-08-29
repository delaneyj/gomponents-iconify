package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSolidTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1754 0q61 0 114 21t94 60t63 91t23 114q0 47-10 84t-28 69t-44 62t-58 61l-424-423q31-31 59-57t60-44t67-28t84-10zm-420 290l422 422L753 1713L0 2046l348-759l986-997z"/>`),
		g.Group(children),
	)
}