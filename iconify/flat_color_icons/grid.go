package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M7 7v34h34V7H7zm32 8h-6V9h6v6zm-14 0V9h6v6h-6zm6 2v6h-6v-6h6zm-8-2h-6V9h6v6zm0 2v6h-6v-6h6zm-8 6H9v-6h6v6zm0 2v6H9v-6h6zm2 0h6v6h-6v-6zm6 8v6h-6v-6h6zm2 0h6v6h-6v-6zm0-2v-6h6v6h-6zm8-6h6v6h-6v-6zm0-2v-6h6v6h-6zM15 9v6H9V9h6zM9 33h6v6H9v-6zm24 6v-6h6v6h-6z"/>`),
		g.Group(children),
	)
}