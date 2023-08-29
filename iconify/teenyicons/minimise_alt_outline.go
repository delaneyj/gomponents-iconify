package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimiseAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13 5.5H9.5m0 0V2m0 3.5l4-4M5.5 13V9.5m0 0H2m3.5 0l-4 4"/>`),
		g.Group(children),
	)
}