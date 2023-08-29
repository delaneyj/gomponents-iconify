package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdGroupOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.7 17.875L18.825 16H20V6H8.825l-2.7-2.7q.2-.6.713-.95T8 2h12q.825 0 1.413.588T22 4v12q0 .65-.35 1.163t-.95.712Zm-.2 5.425L15.2 18H8q-.825 0-1.413-.588T6 16V8.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM8 16h5.2L8 10.8V16Zm-4 6q-.825 0-1.413-.588T2 20V6h2v14h14v2H4Zm6.625-8.575Zm2.85-2.775Z"/>`),
		g.Group(children),
	)
}