package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropRotateTl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 4H7a4 4 0 0 0-4 4v4"/><path d="M7.5 1.5L10 4L7.5 6.5M20 17v-6a1 1 0 0 0-1-1h-6m-5 0h2m10 12v-2M10 8v11a1 1 0 0 0 1 1h11"/></g>`),
		g.Group(children),
	)
}