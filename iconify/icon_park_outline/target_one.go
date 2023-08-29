package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M23.889 39.889c8.84 0 16-7.16 16-16s-7.16-16-16-16s-16 7.16-16 16s7.16 16 16 16Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M23.889 31.889c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8Z"/><path fill="currentColor" d="M23.889 25.889c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="M23.889 7.889v-4m14 40l-4-7m-20 0l-4 7"/></g>`),
		g.Group(children),
	)
}