package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opencalc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.422 27.209V21.79m2.723 2.71h-5.417m.787 11.068l3.83-3.83m.01 3.841l-3.83-3.83m17.751 3.817H27.86m5.42-11.069h-5.417m5.418 7.315h-5.417"/><rect width="21.865" height="7.497" x="13.067" y="10.001" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 4.5h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}