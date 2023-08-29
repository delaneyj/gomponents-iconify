package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-2.925 0-4.963-2.038T1 13q0-2.675 1.763-4.663T7.175 6.05L5.6 4.4L7 3l4 4l-4 4l-1.4-1.425L7.075 8.1Q5.3 8.45 4.15 9.825T3 13q0 2.075 1.462 3.538T8 18h3v2H8Zm5-9V4h9v7h-9Zm0 9v-7h9v7h-9Zm2-2h5v-3h-5v3Z"/>`),
		g.Group(children),
	)
}