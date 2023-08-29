package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 11c0-4.8 4-6 6-6c4.8 0 6 4 6 6v4l2 2H4l2-2v-4z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5c-2 0-6 1.2-6 6v4l-2 2h16l-2-2v-4c0-2-1.2-6-6-6zm0 0V3M9 18c0 1 .6 3 3 3s3-2 3-3"/></g>`),
		g.Group(children),
	)
}