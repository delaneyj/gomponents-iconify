package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdGroupOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L15.2 18H8q-.825 0-1.413-.588T6 16V8.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM4 22q-.825 0-1.413-.588T2 20V6h2v14h14v2H4Zm16.7-4.125L8.825 6H20V4H8v1.175L6.125 3.3q.2-.6.713-.95T8 2h12q.825 0 1.413.588T22 4v12q0 .65-.35 1.163t-.95.712Z"/>`),
		g.Group(children),
	)
}