package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApertureOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.6 15h10.55M5.641 5.631A9 9 0 1 0 18.36 18.369m1.68-2.318A9 9 0 0 0 7.966 3.953m-.571 3.581l2.416 7.438m7.221-10.336L12.18 8.162M9.846 9.857l-1.349.98m12.062 3.673l-8.535-6.201m.233 12.607l2.123-6.533m.984-3.028l.154-.473M3 3l18 18"/>`),
		g.Group(children),
	)
}