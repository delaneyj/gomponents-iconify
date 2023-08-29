package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshSync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 69q70 0 120 50t50 121q0 49-26 91l-31-31q15-28 15-60q0-53-37.5-90.5T171 112v64L85 91l86-86v64zm0 299v-64l85 85l-85 86v-64q-71 0-121-50T0 240q0-49 26-91l32 31q-15 28-15 60q0 53 37.5 90.5T171 368z"/>`),
		g.Group(children),
	)
}