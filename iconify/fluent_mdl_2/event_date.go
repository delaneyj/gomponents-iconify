package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventDate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1792H0V128h384V0h128v128h1024V0h128v128h384zM128 256v256h1792V256h-256v128h-128V256H512v128H384V256H128zm1792 1536V640H128v1152h1792zm-512-896v640h-128v-486q-27 14-62 26t-66 12V960q12 0 31-6t39-15t36-21t22-21v-1h128zm-384 192q0 39-11 70t-31 58t-44 51t-51 46t-51 46t-47 49h235v128H640v-36q0-19-1-38t4-38t10-36q11-27 33-53t50-53t55-51t51-49t39-47t15-47q0-27-19-45t-45-19q-23 0-40 14t-23 37l-125-26q6-33 23-61t44-48t57-32t64-12q40 0 75 15t61 41t41 61t15 75z"/>`),
		g.Group(children),
	)
}