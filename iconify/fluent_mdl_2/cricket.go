package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cricket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 1408q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100q0-53 20-99t55-82t81-55t100-20zm0 384q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10zM1984 0q26 0 45 19t19 45q0 26-19 45l-546 547q26 39 39 84t14 92v27L347 2048L0 1701L1189 512q58 0 105 10t98 43l547-546q19-19 45-19zm-578 807q-4-32-18-60t-36-50t-50-36t-61-19L181 1701l166 166L1406 807z"/>`),
		g.Group(children),
	)
}