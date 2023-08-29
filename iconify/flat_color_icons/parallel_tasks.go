package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParallelTasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M36 13V9H22v13h-9v4h9v13h14v-4H26v-9h10v-4H26v-9z"/><path fill="#D81B60" d="M6 17h10v14H6z"/><path fill="#2196F3" d="M32 6h10v10H32zm0 26h10v10H32zm0-13h10v10H32z"/>`),
		g.Group(children),
	)
}