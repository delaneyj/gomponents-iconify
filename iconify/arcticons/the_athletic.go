package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TheAthletic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.3 6.5h12l12.7 35m4.5 0h-12m-15 0h-12M34 29H14m6.3-18.5L9 41.5"/>`),
		g.Group(children),
	)
}