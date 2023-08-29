package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H128V128h1792zm-119 119H247v1554h1554V247zm-308 1214q0-36-25-60t-65-46t-83-41t-84-49t-64-67t-26-98q0-55 25-94t65-63t88-36t96-11q19 0 41 1t45 3t44 8t40 14v122q-35-26-77-35t-85-10q-23 0-49 4t-50 15t-38 30t-15 48q0 26 10 44t30 33q26 20 56 33t60 27q34 16 68 38t61 49t45 62t17 77q0 60-25 99t-65 63t-90 33t-100 10q-18 0-45-3t-56-8t-55-13t-40-18v-127q17 15 40 27t50 22t52 13t50 5q25 0 52-3t50-15t37-31t15-52zm-726 203q-41 0-79-12t-71-36t-56-56t-37-72l103-32q16 43 51 71t84 29q38 0 61-18t36-47t17-62t5-63V908h123v477q0 56-14 107t-43 89t-73 60t-107 23z"/>`),
		g.Group(children),
	)
}