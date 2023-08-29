package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24.008 41L24 23m16.518 11.316A9.21 9.21 0 0 0 44 24c-1.213-3.83-4.93-5.929-8.947-5.925h-2.321a14.737 14.737 0 1 0-25.31 13.428"/><path d="M30.364 27.636L24 21.272l-6.364 6.364"/></g>`),
		g.Group(children),
	)
}