package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryEduOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-.825 0-1.413-.588T6 18v-2q0-.425.288-.713T7 15h2v-2.25q-.875-.05-1.663-.388T5.9 11.35v-1.1H4.75l-2.4-2.4q-.425-.425-.425-.875t.425-.8q.725-.625 1.787-.95T6.4 4.9q.675 0 1.312.1T9 5.375V5q0-.425.288-.712T10 4h10q.425 0 .713.288T21 5v12q0 1.25-.875 2.125T18 20H8Zm3-5h5q.425 0 .713.288T17 16v1q0 .425.288.713T18 18q.425 0 .713-.288T19 17V6h-8v.6l5.7 5.7q.175.175.238.325T17 13q0 .425-.288.713T16 14q-.225 0-.375-.063T15.3 13.7l-2.55-2.55l-.2.2q-.35.35-.738.625T11 12.4V15ZM5.6 8.25h2.3v2.15q.3.2.625.275t.675.075q.575 0 1.038-.175t.912-.625l.2-.2l-1.4-1.4q-.725-.725-1.625-1.088T6.4 6.9q-.5 0-.95.075t-.9.225L5.6 8.25ZM8 18h7.15q-.075-.225-.113-.475T15 17H8v1Zm0 0v-1v1Z"/>`),
		g.Group(children),
	)
}