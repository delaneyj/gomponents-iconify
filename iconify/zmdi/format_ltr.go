package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 173q-35 0-60-25T0 88t25-60T85 3h171v42h-43v235h-42V45h-43v235H85V173zm256 171l-85 85v-64H0v-42h256v-64z"/>`),
		g.Group(children),
	)
}