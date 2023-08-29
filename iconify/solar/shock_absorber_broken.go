package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShockAbsorberBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14 20a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM6 4c0-.943 0-1.414.293-1.707C6.586 2 7.057 2 8 2h8c.943 0 1.414 0 1.707.293C18 2.586 18 3.057 18 4c0 .943 0 1.414-.293 1.707C17.414 6 16.943 6 16 6H8c-.943 0-1.414 0-1.707-.293C6 5.414 6 4.943 6 4Z"/><path d="M8.5 16.5A1.5 1.5 0 0 1 10 15h4a1.5 1.5 0 0 1 0 3h-4a1.5 1.5 0 0 1-1.5-1.5Zm5.5-1v-10m-4 10V6"/><path stroke-linecap="round" d="m8 8l8 2m-8 1.5l8 2m4-2h2m-18 0H2m17.071 2.8l.707.7m-14.85-.7l-.706.7m14.85-6.3l.706-.7m-14.85.7L4.223 8"/></g>`),
		g.Group(children),
	)
}