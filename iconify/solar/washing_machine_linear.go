package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachineLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6 22v-1m12 1v-1"/><path stroke="currentColor" stroke-width="1.5" d="M3 10c0-3.771 0-5.657 1.172-6.828C5.343 2 7.229 2 11 2h2c3.771 0 5.657 0 6.828 1.172C21 4.343 21 6.229 21 10v3c0 3.771 0 5.657-1.172 6.828C18.657 21 16.771 21 13 21h-2c-3.771 0-5.657 0-6.828-1.172C3 18.657 3 16.771 3 13v-3Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M3 9h18"/><path stroke="currentColor" stroke-width="1.5" d="M15 15a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6.5 5.5h3"/><path fill="currentColor" d="M15.5 5.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g>`),
		g.Group(children),
	)
}