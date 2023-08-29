package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReviewRequestMirroredSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2029 1939q19 19 19 45t-19 45t-45 19q-26 0-45-19l-785-784q-95 80-210 121t-240 42q-97 0-187-25t-168-71t-143-110t-110-142t-71-169T0 704q0-97 25-187t71-168t110-143T348 96t169-71T704 0q97 0 187 25t168 71t143 110t110 142t71 169t25 187q0 124-41 239t-122 211l784 785zM768 1024H640v128h128v-128zm264-512q0-66-29-115t-76-83t-105-49t-118-17q-60 0-118 19t-105 56t-76 89t-29 117q0 53 19 92t47 70t62 55t61 48t48 48t19 54h144q0-53-19-92t-47-70t-62-55t-61-48t-48-48t-19-54q0-34 18-59t45-43t59-26t62-9q27 0 59 6t60 20t46 38t19 56h144z"/>`),
		g.Group(children),
	)
}