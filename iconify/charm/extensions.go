package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extensions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M1.75 8.75h5.5v5.5m5-12.5v4m-2-2h4m-12.5-1v11.5h11.5v-5.5h-6v-6z"/></g>`),
		g.Group(children),
	)
}