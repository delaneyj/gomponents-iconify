package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationMultipleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V7q0-.425.288-.713T2 6q.425 0 .713.288T3 7v12h14q.425 0 .713.288T18 20q0 .425-.288.713T17 21H3Zm10.5-8.65L7 8.55V15h14V8.55l-6.5 3.8q-.225.125-.5.125t-.5-.125ZM7 17q-.825 0-1.413-.588T5 15V6.5q0-.525.238-1t.712-.75L14 0l2.45 1.45l-1.5 1.45l-.95-.55l-6.8 4l6.8 4l6.8-4l-1.35-.8L20.9 4.1l1.15.65q.475.275.712.75t.238 1V15q0 .825-.588 1.413T21 17H7Zm6.85-10.8l3.95-3.95q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L14.55 8.3q-.3.3-.7.3t-.7-.3l-1.4-1.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l.7.7ZM14 15h7H7h7Z"/>`),
		g.Group(children),
	)
}