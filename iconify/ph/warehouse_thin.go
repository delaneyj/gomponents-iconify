package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarehouseThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 188h-12V54.66l12.84-2.75a4 4 0 1 0-1.68-7.82l-224 48A4 4 0 0 0 16 100a4.07 4.07 0 0 0 .84-.09L28 97.52V188H16a4 4 0 0 0 0 8h224a4 4 0 0 0 0-8ZM36 95.81l184-39.43V188h-32v-60a4 4 0 0 0-4-4H72a4 4 0 0 0-4 4v60H36ZM180 156H76v-24h104Zm-104 8h104v24H76Z"/>`),
		g.Group(children),
	)
}