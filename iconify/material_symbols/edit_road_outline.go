package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditRoadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13.05V4h2v7.05l-2 2ZM4 20V4h2v16H4Zm6-12V4h2v4h-2Zm0 6v-4h2v4h-2Zm0 6v-4h2v4h-2Zm12.125-6L20 11.875l.725-.725q.275-.275.7-.275t.7.275l.725.725q.275.275.275.7t-.275.7l-.725.725ZM14 20v-2.125l5.3-5.3l2.125 2.125l-5.3 5.3H14Z"/>`),
		g.Group(children),
	)
}