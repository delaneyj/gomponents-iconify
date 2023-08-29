package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 136h-80a8 8 0 0 0-8 8v57.45a8 8 0 0 0 6.57 7.88l80 14.54a7.61 7.61 0 0 0 1.43.13a8 8 0 0 0 8-8v-72a8 8 0 0 0-8-8Zm-8 70.41l-64-11.63V152h64ZM104 136H40a8 8 0 0 0-8 8v40a8 8 0 0 0 6.57 7.87l64 11.64a8.54 8.54 0 0 0 1.43.13a8 8 0 0 0 8-8V144a8 8 0 0 0-8-8Zm-8 50.05l-48-8.73V152h48ZM221.13 33.86a8 8 0 0 0-6.56-1.73l-80 14.55a8 8 0 0 0-6.57 7.87V112a8 8 0 0 0 8 8h80a8 8 0 0 0 8-8V40a8 8 0 0 0-2.87-6.14ZM208 104h-64V61.22l64-11.63Zm-98.87-49.78a8 8 0 0 0-6.56-1.73l-64 11.64A8 8 0 0 0 32 72v40a8 8 0 0 0 8 8h64a8 8 0 0 0 8-8V60.36a8 8 0 0 0-2.87-6.14ZM96 104H48V78.68L96 70Z"/>`),
		g.Group(children),
	)
}