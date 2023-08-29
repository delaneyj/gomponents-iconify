package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentToAllOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11.825V15q0 .425.288.713T12 16q.425 0 .713-.288T13 15v-3.15l.9.875q.3.275.7.288t.7-.288q.3-.3.3-.712t-.3-.713l-2.6-2.6q-.3-.3-.7-.3t-.7.3l-2.6 2.6q-.3.3-.287.7t.312.7q.3.275.7.288t.7-.288l.875-.875ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}