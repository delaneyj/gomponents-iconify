package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlbumTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 8.5A1.5 1.5 0 0 1 9.5 7h4A1.5 1.5 0 0 1 15 8.5v1a1.5 1.5 0 0 1-1.5 1.5h-4A1.5 1.5 0 0 1 8 9.5v-1ZM9.5 8a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-4ZM2 6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6Zm4 9h10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6v10ZM5 5H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1V5Z"/>`),
		g.Group(children),
	)
}