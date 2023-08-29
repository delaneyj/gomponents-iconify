package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nayapay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.5a9.5 9.5 0 0 0-9.5 9.5v11.04a.5.5 0 0 0 .742.438l1.412-.78A9.605 9.605 0 0 1 21.3 33.5H24a9.5 9.5 0 0 0 0-19Z"/>`),
		g.Group(children),
	)
}