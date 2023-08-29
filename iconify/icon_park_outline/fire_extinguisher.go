package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 20a8 8 0 1 1 16 0v24H6V20Z"/><path d="M30 44h12l-6-18l-6 18Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m36 26l-6 18h12l-6-18Zm0 0v-2c0-7.543 0-11.314-2.343-13.657S27.543 8 20 8h-2m-8 0H6"/><circle cx="14" cy="8" r="4" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/></g>`),
		g.Group(children),
	)
}