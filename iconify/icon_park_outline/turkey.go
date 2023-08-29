package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Turkey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 36h40l-5 8H9l-5-8Zm7-6v6h26v-5c0-6-3-8-3-8c2-2.5 3-6 0-8s-5.5 0-7 2c0 0-16-2-16 13Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M21 24s-1 .5-2 2s-1 3-1 3"/><path stroke="currentColor" stroke-width="4" d="m39 9l-5 6"/><circle cx="38.356" cy="7.483" r="2.5" fill="currentColor" transform="rotate(35.072 38.356 7.483)"/><circle cx="40.811" cy="9.206" r="2.5" fill="currentColor" transform="rotate(35.072 40.81 9.206)"/></g>`),
		g.Group(children),
	)
}