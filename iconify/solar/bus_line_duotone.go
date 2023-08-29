package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4 10c0-3.771 0-5.657 1.172-6.828C6.343 2 8.229 2 12 2c3.771 0 5.657 0 6.828 1.172C20 4.343 20 6.229 20 10v2c0 3.771 0 5.657-1.172 6.828C17.657 20 15.771 20 12 20c-3.771 0-5.657 0-6.828-1.172C4 17.657 4 15.771 4 12v-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 13h16" opacity=".5"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.5 16H17M7 16h1.5"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 19.5V21a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1v-1m8.5-.5V21a1 1 0 0 1-1 1h-1.5a1 1 0 0 1-1-1v-1M20 9h1a1 1 0 0 1 1 1v1a1 1 0 0 1-.4.8L20 13M4 9H3a1 1 0 0 0-1 1v1a1 1 0 0 0 .4.8L4 13" opacity=".5"/><path stroke-linecap="round" d="M19.5 5h-15" opacity=".5"/></g>`),
		g.Group(children),
	)
}