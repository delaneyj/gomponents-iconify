package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedsideTableBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M19 22v-1.5M5 22v-1.5"/><path fill="currentColor" d="M13 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-6a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2 14h12m8 0h-4M2 8h20m-8-6c3.771 0 5.657 0 6.828 1.172C22 4.343 22 6.229 22 10v2c0 3.771 0 5.657-1.172 6.828C19.657 20 17.771 20 14 20h-4c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12v-2c0-3.771 0-5.657 1.172-6.828C4.343 2 6.229 2 10 2"/></g>`),
		g.Group(children),
	)
}