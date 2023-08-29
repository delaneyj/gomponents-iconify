package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LarkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M3.494 17.72L41.678 6.321L30.364 44.59l-8.88-8.88l.041-9.234l-9.546-.27l-8.485-8.486Z"/><path fill="currentColor" fill-rule="evenodd" d="M27.535 14.89a4 4 0 1 0 5.657 5.658a4 4 0 0 0-5.657-5.657Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M41.678 6.406L30.364 17.719"/></g>`),
		g.Group(children),
	)
}