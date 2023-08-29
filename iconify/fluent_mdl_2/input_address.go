package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputAddress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 768v512h1152v128H0V640h1033q15 35 36 67t48 61H128zm1783-128h137v768h-384v-128h256V768h-93q26-29 47-61t37-67zm-375 122v1030l-64 128l-64-128V762q-56-12-103-41t-81-70t-53-94t-19-109q0-66 25-124t68-102t102-69t125-25q66 0 124 25t101 69t69 102t26 124q0 57-19 109t-53 93t-81 71t-103 41zm128-314q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75q0 40 15 75t41 61t61 41t75 15q40 0 75-15t61-41t41-61t15-75z"/>`),
		g.Group(children),
	)
}