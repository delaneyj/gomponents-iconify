package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodPressure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M6 18.724C6 12.641 10.036 7 15.563 7c3.835 0 6.68 2.53 8.437 6.121C25.756 9.531 28.602 7 32.438 7C37.965 7 42 12.642 42 18.724a14.54 14.54 0 0 1-.35 3.155A10.48 10.48 0 0 0 33.5 18C27.701 18 23 22.701 23 28.5a10.5 10.5 0 0 0 5.912 9.447C26.122 39.908 24 41 24 41S6 32.304 6 18.724Z"/><path d="M36 29a2 2 0 1 1-3.97-.348l-2.42-2.42l1.415-1.414l2.299 2.3A2 2 0 0 1 36 29Z"/><path fill-rule="evenodd" d="M42 28.5a8.5 8.5 0 1 1-17 0a8.5 8.5 0 0 1 17 0Zm-2 0a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}