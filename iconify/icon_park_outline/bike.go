package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10.5 42a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17ZM37 42a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18.997 6h9L37 33"/><path d="m11.057 33l20.625-16.237L11.057 33Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m11.057 33l20.625-16.237m0-1.763h8.472L42 10M8 15.974h7M15 16l3.273 10.421"/></g>`),
		g.Group(children),
	)
}