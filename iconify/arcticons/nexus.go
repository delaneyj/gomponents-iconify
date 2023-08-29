package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nexus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.526 25.5H25.5v3.024L39.475 42.5l3.025-3.025L28.526 25.5zM22.5 28.526V25.5h-3.024L5.5 39.475L8.525 42.5L22.5 28.526zM19.474 22.5H22.5v-3.024L8.525 5.5L5.5 8.525L19.474 22.5zm6.026-3.026V22.5h3.024L42.5 8.525L39.475 5.5L25.5 19.474z"/>`),
		g.Group(children),
	)
}