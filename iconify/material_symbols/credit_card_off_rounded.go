package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.775 18.925L14.85 12H20V8h-9.15l-4-4H20q.825 0 1.413.588T22 6v12q0 .25-.05.488t-.175.437ZM4 12h5.15l-4-4H4v4Zm15.75 10.575L17.15 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H3.15l-1.8-1.8q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Z"/>`),
		g.Group(children),
	)
}