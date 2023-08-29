package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenCast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 1792q0-27-10-50t-27-40t-41-28t-50-10v-128q53 0 99 20t82 55t55 81t20 100h-128zM128 384v896h1024v128h-128v128h128v128H640v-128h256v-128H0V256h1920v866q-59-56-128-102V384H128zm1536 1408q0-79-30-149t-82-122t-123-83t-149-30v-128q106 0 199 40t163 109t110 163t40 200h-128zm256 0q0-132-50-248t-138-204t-203-137t-249-51v-128q106 0 204 27t183 78t156 120t120 155t77 184t28 204h-128z"/>`),
		g.Group(children),
	)
}