package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectingAirports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.6 21.4L13 17h-3q-.425 0-.713-.288T9 16q0-.425.288-.713T10 15h3l2.6-4.4h1.1L15.4 15h2.85l.75-1h1l-.6 2l.6 2h-1l-.75-1H15.4l1.3 4.4h-1.1Zm-8.3-8L8.6 9H5.75L5 10H4l.6-2L4 6h1l.75 1H8.6L7.3 2.6h1.1L11 7h3q.425 0 .713.288T15 8q0 .425-.288.713T14 9h-3l-2.6 4.4H7.3Z"/>`),
		g.Group(children),
	)
}