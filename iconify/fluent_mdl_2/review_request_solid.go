package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReviewRequestSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2029 1939q19 19 19 45t-19 45t-45 19q-26 0-45-19l-785-784q-95 80-210 121t-240 42q-97 0-187-25t-168-71t-143-110t-110-142t-71-169T0 704q0-97 25-187t71-168t110-143T348 96t169-71T704 0q97 0 187 25t168 71t143 110t110 142t71 169t25 187q0 124-41 239t-122 211l784 785zM768 1024H640v128h128v-128zm8-128q0-30 19-54t47-47t62-48t61-56t48-70t19-92q0-65-29-117t-75-88t-105-56t-119-20q-59 0-117 16t-106 50t-76 82t-29 116h144q0-33 18-56t47-37t60-21t59-6q29 0 61 9t60 26t45 42t18 60q0 30-19 54t-47 47t-62 48t-61 56t-48 70t-19 92h144z"/>`),
		g.Group(children),
	)
}