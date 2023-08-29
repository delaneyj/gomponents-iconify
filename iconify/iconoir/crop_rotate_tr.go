package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropRotateTr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 10V7a4 4 0 0 0-4-4h-4"/><path d="M22.5 7.5L20 10l-2.5-2.5M14 17v-6a1 1 0 0 0-1-1H7m-5 0h2m10 12v-2M4 8v11a1 1 0 0 0 1 1h11"/></g>`),
		g.Group(children),
	)
}