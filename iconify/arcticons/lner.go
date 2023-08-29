package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.672 29.153V18.847h3.349c1.933 0 3.479 1.546 3.479 3.478s-1.546 3.48-3.48 3.48h-3.348m3.475-.004l3.224 3.224M8.5 18.847v10.306h5.153m2.346 0V18.847l6.829 10.306V18.847M25.172 24h3.35m1.804 5.153h-5.155V18.847h5.154"/>`),
		g.Group(children),
	)
}