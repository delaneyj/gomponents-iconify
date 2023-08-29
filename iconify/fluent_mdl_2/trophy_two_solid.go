package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrophyTwoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 1536q0-62 29-109t76-80t103-50t112-17q55 0 111 17t103 49t76 80t30 110H640zm864 128q13 0 22 9t10 23v224H384v-224q0-13 9-22t23-10h1088zm288-1408q27 0 50 10t40 27t28 41t10 50v192q0 98-36 180t-98 141t-147 93t-180 34q-41 62-96 110t-121 80t-137 49t-145 17q-74 0-145-16t-137-49t-120-81t-97-110q-96 0-180-33t-146-93t-99-142T0 576V384q0-27 10-50t27-40t41-28t50-10h256V128h1152v128h256zM384 384H128v192q0 57 19 109t53 93t81 71t103 41V384zm1408 0h-256v506q56-12 103-41t81-70t53-94t19-109V384z"/>`),
		g.Group(children),
	)
}