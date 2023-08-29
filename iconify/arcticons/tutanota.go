package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tutanota(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33.18a1 1 0 0 0 1.28 1c8-2.34 20.44-5.85 23.38-10.15c1.57-3-8.56-5.07-8.14-8.21s6.46-4.3 9-4.83a8 8 0 0 1 2.92 0s-8.35 2.6-8.13 4.4s9.77 2.32 14.89 4.74l1.13.56a.49.49 0 0 0 .72-.44V7.5a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}