package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 16.95l-1.45-1.4l.65-.65V9.1L14.9 5H9.1l-.65.65l-1.4-1.45L8.25 3h7.5L21 8.25v7.45l-1.2 1.25ZM13 10.2V7h-2v1.2l2 2Zm7.5 13.1l-3.55-3.55l-1.2 1.25h-7.5L3 15.7V8.25l1.2-1.2L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4Zm-7.1-12.7ZM9.1 19h5.8l.65-.65l-9.9-9.9L5 9.1v5.8L9.1 19Zm2.9-2q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Zm-1.4-3.6Z"/>`),
		g.Group(children),
	)
}