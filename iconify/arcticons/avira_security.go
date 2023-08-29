package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AviraSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.701 13.096c-7.595-7.595-19.91-7.595-27.505 0c-7.595 7.595-7.595 19.91 0 27.505l27.505-27.505Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.949 26.848l11.066 11.066a3.799 3.799 0 0 0 5.372 0h0a3.799 3.799 0 0 0 0-5.372l-1.44-1.44"/>`),
		g.Group(children),
	)
}