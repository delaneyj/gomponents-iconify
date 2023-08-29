package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DepthLab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.887 14.25L24 4.5L7.113 14.25v19.5L24 43.5l16.887-9.75v-19.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.113 14.25L24 24l16.887-9.75M24 43.5V24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.12 11.668v16.443l-14.24 8.221V19.889l14.24-8.221z"/>`),
		g.Group(children),
	)
}