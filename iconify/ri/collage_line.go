package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollageLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3.106a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-16a1 1 0 0 1 1-1h16Zm-8.811 10.159L5 14.355v4.751h7.218l-1.03-5.841ZM19 5.106h-7.219l2.468 14H19v-14Zm-9.25 0H5v7.218l5.841-1.029L9.75 5.106Z"/>`),
		g.Group(children),
	)
}