package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Metrodroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.55 8.4H6.45a2 2 0 0 0-1.95 1.95v27.3a2 2 0 0 0 2 2h35.1a2 2 0 0 0 2-2v-27.3a2 2 0 0 0-2.05-1.95Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m37 21.75l2.92 3.12a.61.61 0 0 1 .06.86h0l-2.92 3.12M7.87 29.88V16.11l6.88 13.79l6.87-13.77V29.9m10.26-13.76v13.73"/><rect width="6.87" height="9.13" x="25.02" y="20.74" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.43"/>`),
		g.Group(children),
	)
}