package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adoberush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.94 33.89V14h6.65a6.67 6.67 0 0 1 0 13.33H9.94m6.65.04l6.58 6.59m14.89-5.05a5 5 0 0 1-5 5h0a5 5 0 0 1-5-5v-8m10 12.98V20.94"/>`),
		g.Group(children),
	)
}