package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteroidBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M10 19a6 6 0 0 0-7.915-5.688M17 8a6.002 6.002 0 0 0 4.823 5.885m-3.505-9.637a5.95 5.95 0 0 0-.515.752"/><path d="M16 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke-linecap="round" d="M13 8.5a2.5 2.5 0 1 0-2.5 2.5"/><path stroke-linecap="round" d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}