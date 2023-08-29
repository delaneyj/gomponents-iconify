package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Penup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.036 35.23h14.103L25.386 6.3a1.602 1.602 0 0 0-2.772 0L5.86 35.23h11.06l7.27 7.27V24.34m-4.155-13.474h7.93"/>`),
		g.Group(children),
	)
}