package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasMeterOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-1.65 0-2.825-1.175T4 18V8q0-1.65 1.175-2.825T8 4h1V2h2v2h2V2h2v2h1q1.65 0 2.825 1.175T20 8v10q0 1.65-1.175 2.825T16 22H8Zm0-2h8q.825 0 1.413-.588T18 18V8q0-.825-.588-1.413T16 6H8q-.825 0-1.413.588T6 8v10q0 .825.588 1.413T8 20Zm0-10h8V8H8v2Zm4 8q1.05 0 1.775-.713t.725-1.737q0-.825-.475-1.413T12 11.75q-1.575 1.8-2.038 2.4t-.462 1.4q0 1.025.725 1.738T12 18ZM6 6v14V6Z"/>`),
		g.Group(children),
	)
}