package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRightDoubleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10v9h2v-7h6.172l-3.95 3.95l1.414 1.414L16 11L9.636 4.636L8.222 6.05l3.95 3.95H4Zm11.25-5.364L13.837 6.05l4.95 4.95l-4.95 4.95l1.415 1.414L21.615 11L15.25 4.636Z"/>`),
		g.Group(children),
	)
}