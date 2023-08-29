package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReviewResponseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2029 1939q19 19 19 45t-19 45t-45 19q-26 0-45-19l-785-784q-95 80-210 121t-240 42q-97 0-187-25t-168-71t-143-110t-110-142t-71-169T0 704q0-97 25-187t71-168t110-143T348 96t169-71T704 0q97 0 187 25t168 71t143 110t110 142t71 169t25 187q0 124-41 239t-122 211l784 785zM1146 525l-135-135l-435 434l-156-156l-136 136l292 292l570-571z"/>`),
		g.Group(children),
	)
}