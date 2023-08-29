package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11.9V4h2v5.9l-2 2ZM4 20V4h2v16H4Zm6-12V4h2v4h-2Zm0 6v-4h2v4h-2Zm0 6v-4h2v4h-2Zm9.075-7.225l2.15 2.1l-5.1 5.125H14v-2.125l5.075-5.1Zm2.875 1.4l-2.125-2.125l.7-.7q.3-.3.725-.3t.7.3l.7.725q.275.3.275.713t-.275.687l-.7.7Z"/>`),
		g.Group(children),
	)
}