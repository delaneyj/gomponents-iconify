package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 24v4H6v-4H4v4l.008-.005A1.998 1.998 0 0 0 6 30h20a2 2 0 0 0 2-2v-4zM6 12l1.411 1.405L15 5.825V24h2V5.825l7.591 7.58L26 12L16 2L6 12z"/>`),
		g.Group(children),
	)
}