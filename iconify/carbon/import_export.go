package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 24v4H4v-4H2v4l.008-.005A1.998 1.998 0 0 0 4 30h24a2 2 0 0 0 2-2v-4zm-.4-9.4L24 18.2V4h-2v14.2l-3.6-3.6L17 16l6 6l6-6l-1.4-1.4zM9 4l-6 6l1.4 1.4L8 7.8V22h2V7.8l3.6 3.6L15 10L9 4z"/>`),
		g.Group(children),
	)
}