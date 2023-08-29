package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinorCrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 6.4l-3-3L6.4 2l3 3L8 6.4Zm8 0L14.6 5l3-3L19 3.4l-3 3ZM11 5V0h2v5h-2ZM4 24q-.425 0-.713-.288T3 23v-8l2.1-6q.15-.45.537-.725T6.5 8h11q.475 0 .863.275T18.9 9l2.1 6v8q0 .425-.287.713T20 24h-1q-.425 0-.713-.288T18 23v-1H6v1q0 .425-.288.713T5 24H4Zm1.8-11h12.4l-1.05-3H6.85L5.8 13Zm1.7 6q.625 0 1.063-.438T9 17.5q0-.625-.438-1.063T7.5 16q-.625 0-1.063.438T6 17.5q0 .625.438 1.063T7.5 19Zm9 0q.625 0 1.063-.438T18 17.5q0-.625-.438-1.063T16.5 16q-.625 0-1.063.438T15 17.5q0 .625.438 1.063T16.5 19Z"/>`),
		g.Group(children),
	)
}