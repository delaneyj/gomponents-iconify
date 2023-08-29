package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimiseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M9.5 9.5H13m-3.5 0V13m0-3.5l4 4m-.5-8H9.5m0 0V2m0 3.5l4-4M2 5.5h3.5m0 0V2m0 3.5l-4-4m4 11.5V9.5m0 0H2m3.5 0l-4 4"/>`),
		g.Group(children),
	)
}