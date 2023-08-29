package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bspb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="17.59" r="13.06" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="24" cy="36.21" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="13.06" ry="3.64"/><ellipse cx="24" cy="42.66" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="13.06" ry=".87"/>`),
		g.Group(children),
	)
}