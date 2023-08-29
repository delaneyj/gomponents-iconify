package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.43 5.5A1.92 1.92 0 0 0 5.5 7.43v33.14a1.92 1.92 0 0 0 1.93 1.93h33.14a1.92 1.92 0 0 0 1.93-1.93V7.43a1.92 1.92 0 0 0-1.93-1.93Zm3.13 17.41l12.58 7.27l-12.58 7.27m27.33 0H24.44"/>`),
		g.Group(children),
	)
}