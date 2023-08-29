package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Offerup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.369 28.832l30.755-5.516l5.376-9.143l-8.245-5.958L4.5 13.73l2.869 15.102z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.494 28.272l7.997 11.513l23.522-15.73l1.487-9.882"/><circle cx="39.339" cy="14.912" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}