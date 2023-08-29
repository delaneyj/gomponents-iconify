package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M3 5v16"/><path d="M3 16h18M3 7h18m-3 9V7m-4 9V7m-4 9V7m-4 9V7M3 19h18"/><path stroke-linecap="round" d="M21 5v16m0-16a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM3 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`),
		g.Group(children),
	)
}