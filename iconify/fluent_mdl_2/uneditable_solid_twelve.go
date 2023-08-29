package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UneditableSolidTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1908 562l-424-423q31-31 59-57t60-44t67-28t84-10q61 0 114 21t94 60t63 91t23 114q0 47-10 84t-28 70t-44 61t-58 61zm-152 150l-282 282l-420-420l280-284l422 422zM25 146L146 25l1877 1877l-121 121l-729-729l-420 419L0 2046l348-759l407-411L25 146z"/>`),
		g.Group(children),
	)
}