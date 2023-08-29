package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropRotateBl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 14v3a4 4 0 0 0 4 4h4"/><path d="M1.5 16.5L4 14l2.5 2.5M20 11V5a1 1 0 0 0-1-1h-6M8 4h2m10 12v-2M10 2v11a1 1 0 0 0 1 1h11"/></g>`),
		g.Group(children),
	)
}