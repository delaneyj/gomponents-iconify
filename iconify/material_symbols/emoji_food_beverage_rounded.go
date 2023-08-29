package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiFoodBeverageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20q0-.425.288-.713T5 19h14q.425 0 .713.288T20 20q0 .425-.288.713T19 21H5ZM18 8h2V5h-2v3ZM8 17q-1.65 0-2.825-1.175T4 13V5.225q0-.925.65-1.575T6.225 3H9v2.4L7.2 6.85q-.05.05-.2.4v4.25q0 .2.15.35t.35.15h4q.2 0 .35-.15t.15-.35V7.25q0-.05-.2-.4L10 5.4V3h10q.825 0 1.413.587T22 5v3q0 .825-.588 1.413T20 10h-2v3q0 1.65-1.175 2.825T14 17H8Z"/>`),
		g.Group(children),
	)
}