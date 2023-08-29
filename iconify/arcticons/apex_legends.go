package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApexLegends(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 20.039l8.714 15.264H26.75l10.33 7.088l5.42-4.378L24 5.609L5.5 38.013l5.42 4.378l10.33-7.088h-5.964L24 20.039z"/>`),
		g.Group(children),
	)
}