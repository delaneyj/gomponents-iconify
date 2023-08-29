package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shpock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="20.24" cy="27.76" r="14.74" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34.98" cy="13.021" r="7.52" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}