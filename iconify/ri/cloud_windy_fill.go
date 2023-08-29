package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudWindyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 19v-3.993H2.074a8 8 0 1 1 14.383-6.908A5.5 5.5 0 1 1 17.5 19h-3.499Zm-8 2h10v2H6v-2Zm-4-4h10v2H2v-2Z"/>`),
		g.Group(children),
	)
}