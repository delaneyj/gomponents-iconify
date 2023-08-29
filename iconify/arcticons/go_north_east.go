package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoNorthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.973 13.425a12.573 12.573 0 0 0-6.084-1.556c-6.842 0-12.389 5.413-12.389 12.089s5.547 12.088 12.39 12.088c6.073 0 11.126-4.264 12.187-9.89h-6.229m6.393-5.892a9.135 9.135 0 0 1 5.255-1.646c4.973 0 9.004 3.921 9.004 8.757s-4.032 8.756-9.005 8.756a9.185 9.185 0 0 1-3.741-.79"/>`),
		g.Group(children),
	)
}