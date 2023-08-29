package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearNight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1536q141 0 272-36t244-104t207-160t161-207t103-245t37-272q0-133-34-261T1019 8q128 16 245 61t217 115t181 161t138 199t89 229t31 251q0 141-36 272t-104 244t-160 207t-207 161t-245 103t-272 37q-138 0-269-36t-246-103t-212-164T5 1528q31 4 61 6t62 2zm768 384q124 0 238-32t214-90t181-140t140-181t91-214t32-239q0-136-40-263t-112-236t-176-194t-229-136q45 155 45 317q0 146-35 282t-100 258t-157 225t-205 182t-244 129t-277 68q128 128 290 196t344 68z"/>`),
		g.Group(children),
	)
}