package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultipleStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 20l-4-4l4-4l1.425 1.4l-1.6 1.6H11v2H6.825L8.4 18.6L7 20Zm7-3q-.425 0-.713-.288T13 16q0-.425.288-.713T14 15q.425 0 .713.288T15 16q0 .425-.288.713T14 17Zm4 0q-.425 0-.713-.288T17 16q0-.425.288-.713T18 15q.425 0 .713.288T19 16q0 .425-.288.713T18 17Zm-1-5l-1.425-1.4l1.6-1.6H13V7h4.175L15.6 5.4L17 4l4 4l-4 4ZM6 9q-.425 0-.713-.288T5 8q0-.425.288-.713T6 7q.425 0 .713.288T7 8q0 .425-.288.713T6 9Zm4 0q-.425 0-.713-.288T9 8q0-.425.288-.713T10 7q.425 0 .713.288T11 8q0 .425-.288.713T10 9Z"/>`),
		g.Group(children),
	)
}