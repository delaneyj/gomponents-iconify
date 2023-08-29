package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.243 18.506l.01-2.514l2.504.01m13 2.504l-.01-2.514l-2.504.01m-13 13.492l.01 2.514l2.504-.01m13-2.505l-.01 2.515l-2.504-.01M24.016 22l-.032 4M26 24.016l-4-.032"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 43.5h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}