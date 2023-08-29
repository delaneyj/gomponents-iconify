package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalconyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12q-.425 0-.713-.288T8 11q0-.425.288-.713T9 10q.425 0 .713.288T10 11q0 .425-.288.713T9 12Zm6 0q-.425 0-.713-.288T14 11q0-.425.288-.713T15 10q.425 0 .713.288T16 11q0 .425-.288.713T15 12ZM5 22q-.825 0-1.413-.588T3 20v-4q0-.5.238-.963T4 14.275V10q0-1.65.625-3.113t1.713-2.55q1.087-1.087 2.55-1.712T12 2q1.65 0 3.113.625t2.55 1.713q1.087 1.087 1.712 2.55T20 10v4.275q.525.3.763.763T21 16v4q0 .825-.588 1.413T19 22H5Zm0-6v4h2v-4H5Zm4 4h2v-4H9v4Zm-3-6h5V4.075q-2.15.35-3.575 2.012T6 10v4Zm7 0h5v-4q0-2.25-1.425-3.913T13 4.075V14Zm0 6h2v-4h-2v4Zm4 0h2v-4h-2v4Z"/>`),
		g.Group(children),
	)
}