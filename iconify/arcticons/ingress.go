package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ingress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.978 14.75v18.5L24 42.5l16.021-9.25v-18.5L24 5.5L7.978 14.75zM24 5.5v8.022m-8.978 15.889L7.978 33.25m24.999-3.839l7.044 3.839"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 37.269L35.668 17.06H12.332L24 37.269zM12.332 17.06L24 23.796m11.668-6.736L24 23.796m0 13.473V23.796"/>`),
		g.Group(children),
	)
}