package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlusBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 3h340q17 0 30.5 17T427 56v330q0 18-12.5 30.5T384 429H43q-18 0-30.5-12.5T0 387V46q0-18 12.5-30.5T43 3zm108 104q-45 0-76.5 32T43 216t31.5 77t76.5 32q47 0 75.5-29.5T255 219q0-13-2-19H151v38h62q-3 17-18 31.5T151 284q-28 0-47.5-20T84 216t19.5-48t47.5-20q27 0 43 16l30-28q-29-29-73-29zm171 62v31h-31v31h31v31h31v-31h30l1-31h-31v-31h-31z"/>`),
		g.Group(children),
	)
}