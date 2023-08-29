package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerPlugOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 21q-.425 0-.713-.288T9.5 20v-2l-2.925-2.925q-.275-.275-.425-.637T6 13.675V9q0-.6.275-1.125t.8-.8L9 9H8v4.65l3.5 3.5V19h1v-1.85l.925-.925L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-4.25-4.25l-.35.35v2q0 .425-.287.713T13.5 21h-3ZM18 9v3.625q0 .4-.138.763t-.412.637l-.3.275L16 13.15V9h-4.15L8 5.15V4q0-.425.288-.713T9 3q.425 0 .713.288T10 4v3h4V4q0-.425.288-.713T15 3q.425 0 .713.288T16 4v4l-1-1h1q.825 0 1.413.588T18 9Zm-4.05 2.1Zm-3.25 2.425Z"/>`),
		g.Group(children),
	)
}