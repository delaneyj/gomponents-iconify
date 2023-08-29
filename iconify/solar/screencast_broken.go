package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14 20c3.771 0 5.657 0 6.828-1.172C22 17.657 22 15.771 22 12v-1M2 8.5c0-.464 0-.697.02-.892a4 4 0 0 1 3.588-3.589C5.803 4 6.036 4 6.5 4H14c3.771 0 5.657 0 6.828 1.172c.47.47.751 1.054.92 1.828"/><path d="M11 20a9 9 0 0 0-9-9"/><path d="M8 20a6 6 0 0 0-6-6m3 6a3 3 0 0 0-3-3"/></g>`),
		g.Group(children),
	)
}