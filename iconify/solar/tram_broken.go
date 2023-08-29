package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TramBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 10v2c0 3.771 0 5.657-1.172 6.828C17.657 20 15.771 20 12 20c-3.771 0-5.657 0-6.828-1.172C4 17.657 4 15.771 4 12v-2c0-3.771 0-5.657 1.172-6.828C6.343 2 8.229 2 12 2c3.771 0 5.657 0 6.828 1.172c.654.653.943 1.528 1.07 2.828"/><path stroke-linejoin="round" d="M15.5 16H17M7 16h1.5M7 20l-1 2m11-2l1 2M10 2v.5a2 2 0 1 0 4 0V2m6 11h-4M4 13h8"/></g>`),
		g.Group(children),
	)
}