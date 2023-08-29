package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseRemind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 38s-6-5-6-19c0-6.627-5.373-12-12-12c-2.54 0-4.896.79-6.835 2.136M30 38H6s5.57-4.641 5.976-17.5"/><path stroke="currentColor" stroke-width="4" d="M18 38h12a6 6 0 0 1-12 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M24 2a4 4 0 0 0-4 4h8a4 4 0 0 0-4-4Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m7 6.5l34 38"/></g>`),
		g.Group(children),
	)
}