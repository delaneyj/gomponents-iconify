package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 3v6H7v6h6v14h6V15h6V9h-6V3zm2 2h2v6h6v2h-6v14h-2V13H9v-2h6z"/>`),
		g.Group(children),
	)
}