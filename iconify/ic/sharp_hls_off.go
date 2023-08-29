package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpHlsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.83 15h2.67v-3.5H17v-1h2v.5h1.5V9h-5v3.5H19v1h-2V13h-1.17l2 2zM8 10.83V15H6.5v-2.5h-2V15H3V9h1.5v2h2V9.33L1.39 4.22L2.8 2.81l18.38 18.38l-1.41 1.41l-7.6-7.6H10v-2.17l-2-2z"/>`),
		g.Group(children),
	)
}