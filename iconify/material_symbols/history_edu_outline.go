package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryEduOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20q-.825 0-1.413-.588T6 18v-3h3v-2.25q-.875-.05-1.663-.388T5.9 11.35v-1.1H4.75L1.5 7q.9-1.15 2.225-1.625T6.4 4.9q.675 0 1.313.1T9 5.375V4h12v13q0 1.25-.875 2.125T18 20H8Zm3-5h6v2q0 .425.288.713T18 18q.425 0 .713-.288T19 17V6h-8v.6l6 6V14h-1.4l-2.85-2.85l-.2.2q-.35.35-.738.625T11 12.4V15ZM5.6 8.25h2.3v2.15q.3.2.625.275t.675.075q.575 0 1.038-.175t.912-.625l.2-.2l-1.4-1.4q-.725-.725-1.625-1.088T6.4 6.9q-.5 0-.95.075t-.9.225L5.6 8.25ZM15 17H8v1h7.15q-.075-.225-.113-.475T15 17Zm-7 1v-1v1Z"/>`),
		g.Group(children),
	)
}