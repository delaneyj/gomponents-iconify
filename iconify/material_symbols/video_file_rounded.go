package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFileRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V20q0 .825-.588 1.413T18 22H6Zm7-14q0 .425.288.713T14 9h4l-5-5v4ZM9 18h4q.425 0 .713-.288T14 17v-1l1.275.675q.25.125.488-.025t.237-.425v-2.45q0-.275-.238-.425t-.487-.025L14 14v-1q0-.425-.288-.712T13 12H9q-.425 0-.713.288T8 13v4q0 .425.288.713T9 18Z"/>`),
		g.Group(children),
	)
}