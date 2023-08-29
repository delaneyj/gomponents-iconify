package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoMeetingRoomOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 16.15l-2-2V6h-2v6.15l-2-2V5H7.85l-2-2H14q.4 0 .563.363T15 4h3q.425 0 .713.288T19 5v11.15Zm.1 5.75L15 17.8V20q0 .425-.288.713T14 21H4q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h1V7.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM7 19h6v-3.2l-6-6V19Zm3.425-11.425ZM10 12.8Z"/>`),
		g.Group(children),
	)
}