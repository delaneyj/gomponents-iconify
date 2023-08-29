package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TranslateRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.1 18.95l-.875 2.425q-.1.275-.35.45t-.55.175q-.5 0-.812-.413t-.113-.912l3.8-10.05q.125-.275.375-.45t.55-.175h.75q.3 0 .55.175t.375.45L22.6 20.7q.2.475-.1.888t-.8.412q-.325 0-.562-.188t-.363-.487l-.825-2.375H15.1Zm.6-1.75h3.6l-1.75-4.95h-.1L15.7 17.2ZM9 14l-4.3 4.3q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l4.35-4.35q-.95-1.05-1.663-2.175T4.75 8h2.1q.45.9.963 1.625T9.05 11.15q1.1-1.2 1.825-2.463T12.1 6H2q-.425 0-.713-.288T1 5q0-.425.288-.713T2 4h6V3q0-.425.288-.713T9 2q.425 0 .713.288T10 3v1h6q.425 0 .713.288T17 5q0 .425-.288.713T16 6h-1.9q-.525 1.775-1.425 3.45T10.45 12.6l2.4 2.45l-.75 2.05L9 14Z"/>`),
		g.Group(children),
	)
}