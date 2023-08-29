package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagsafeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 2.5v-2h8v2m-6 9V15m4-3.5V15m-5-5.5v2h6v-2m-9-7h12v7h-12v-7Z"/>`),
		g.Group(children),
	)
}