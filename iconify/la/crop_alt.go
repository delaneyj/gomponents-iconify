package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 4v4H4v2h4v14h14v4h2v-4h4v-2H10V4H8zm4 4v2h10v10h2V8H12z"/>`),
		g.Group(children),
	)
}