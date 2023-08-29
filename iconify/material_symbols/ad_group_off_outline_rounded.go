package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdGroupOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.7 17.875L18.825 16H20V6H8.825l-2.7-2.7q.2-.6.713-.95T8 2h12q.825 0 1.413.588T22 4v12q0 .65-.35 1.163t-.95.712ZM8 18q-.825 0-1.412-.588T6 16V8.8L1.4 4.2q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L15.2 18H8Zm0-2h5.2L8 10.8V16Zm-4 6q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4Zm6.625-8.575Zm2.85-2.775Z"/>`),
		g.Group(children),
	)
}