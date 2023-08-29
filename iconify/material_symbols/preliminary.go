package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preliminary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17q.425 0 .713-.288T10 16q0-.425-.288-.713T9 15q-.425 0-.713.288T8 16q0 .425.288.713T9 17Zm3 0q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm3 0q.425 0 .713-.288T16 16q0-.425-.288-.713T15 15q-.425 0-.713.288T14 16q0 .425.288.713T15 17Zm-4.05-3l5.625-5.65l-1.4-1.425l-4.25 4.25L8.8 9.05l-1.4 1.4L10.95 14ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}