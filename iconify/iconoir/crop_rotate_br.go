package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropRotateBr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 20h3a4 4 0 0 0 4-4v-4"/><path d="M16.5 22.5L14 20l2.5-2.5M14 11V5a1 1 0 0 0-1-1H7M2 4h2m10 12v-2M4 2v11a1 1 0 0 0 1 1h11"/></g>`),
		g.Group(children),
	)
}