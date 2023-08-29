package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 7.2A4.2 4.2 0 0 1 10.2 3h11.6A4.2 4.2 0 0 1 26 7.2V28a1 1 0 0 1-1.585.81L16 22.735L7.585 28.81A1 1 0 0 1 6 28V7.2ZM10.2 5A2.2 2.2 0 0 0 8 7.2v18.844l7.415-5.355a1 1 0 0 1 1.17 0L24 26.044V7.2A2.2 2.2 0 0 0 21.8 5H10.2Z"/>`),
		g.Group(children),
	)
}