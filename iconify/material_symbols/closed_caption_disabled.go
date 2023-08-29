package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-3.3-3.3H5q-.825 0-1.413-.588T3 18V5.825L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM21 18.125l-3.35-3.35q.175-.125.263-.338T18 14v-1h-1.5v.5h-.125L14.5 11.625V10.5h2v.5H18v-1q0-.425-.288-.713T17 9h-3q-.425 0-.713.288T13 10v.125L6.875 4H19q.825 0 1.413.588T21 6v12.125ZM7 15h3q.425 0 .713-.288T11 14v-.175L10.175 13H9.5v.5h-2v-3.175L6.375 9.2q-.175.125-.275.338T6 10v4q0 .425.288.713T7 15Z"/>`),
		g.Group(children),
	)
}