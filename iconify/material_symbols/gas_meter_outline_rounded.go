package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasMeterOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-1.65 0-2.825-1.175T4 18V8q0-1.65 1.175-2.825T8 4h1V3q0-.425.288-.713T10 2q.425 0 .713.288T11 3v1h2V3q0-.425.288-.713T14 2q.425 0 .713.288T15 3v1h1q1.65 0 2.825 1.175T20 8v10q0 1.65-1.175 2.825T16 22H8Zm0-2h8q.825 0 1.413-.588T18 18V8q0-.825-.588-1.413T16 6H8q-.825 0-1.413.588T6 8v10q0 .825.588 1.413T8 20Zm1-10h6q.425 0 .713-.288T16 9q0-.425-.288-.713T15 8H9q-.425 0-.713.288T8 9q0 .425.288.713T9 10Zm3 8q1.05 0 1.775-.713t.725-1.737q0-.65-.288-1.125T13.05 12.95l-.675-.775Q12.225 12 12 12t-.375.175l-.675.775q-.875 1.025-1.163 1.488T9.5 15.55q0 1.025.725 1.738T12 18ZM6 6v14V6Z"/>`),
		g.Group(children),
	)
}