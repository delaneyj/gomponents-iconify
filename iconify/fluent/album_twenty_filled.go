package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbumTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6a2 2 0 0 1 2-2v12a2 2 0 0 1-2-2V6Zm7 2.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-1ZM5 16V4h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5Zm4.5-9A1.5 1.5 0 0 0 8 8.5v1A1.5 1.5 0 0 0 9.5 11h4A1.5 1.5 0 0 0 15 9.5v-1A1.5 1.5 0 0 0 13.5 7h-4Z"/>`),
		g.Group(children),
	)
}