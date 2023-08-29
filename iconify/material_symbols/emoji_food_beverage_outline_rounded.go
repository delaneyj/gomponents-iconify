package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiFoodBeverageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20q0-.425.288-.713T5 19h14q.425 0 .713.288T20 20q0 .425-.288.713T19 21H5Zm3-4q-1.65 0-2.825-1.175T4 13V5.225q0-.925.65-1.575T6.225 3H20q.825 0 1.413.588T22 5v3q0 .825-.588 1.413T20 10h-2v3q0 1.65-1.175 2.825T14 17H8ZM8 5h8H6h2Zm10 3h2V5h-2v3Zm-4 7q.825 0 1.413-.588T16 13V5h-6v.4l1.8 1.45q.05.05.2.4v4.25q0 .2-.15.35t-.35.15h-4q-.2 0-.35-.15T7 11.5V7.25q0-.05.2-.4L9 5.4V5H6v8q0 .825.588 1.413T8 15h6ZM9 5h1h-1Z"/>`),
		g.Group(children),
	)
}