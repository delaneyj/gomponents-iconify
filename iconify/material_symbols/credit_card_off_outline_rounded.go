package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.775 18.925L20 17.15V12h-5.15l-4-4H20V6H8.85l-2-2H20q.825 0 1.413.588T22 6v12q0 .25-.05.488t-.175.437ZM9.625 12.45Zm4.8-.875ZM9.15 12H4v6h11.15l-6-6ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H4v2h1.15l-3.8-3.8q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L17.15 20H4Z"/>`),
		g.Group(children),
	)
}