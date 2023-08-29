package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoCrashOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6.35l-2.825-2.8l1.4-1.425L12 3.55L15.55 0l1.4 1.4L12 6.35ZM4 24q-.425 0-.713-.288T3 23v-8l2.1-6q.15-.45.537-.725T6.5 8h11q.475 0 .863.275T18.9 9l2.1 6v8q0 .425-.287.713T20 24h-1q-.425 0-.713-.288T18 23v-1H6v1q0 .425-.288.713T5 24H4Zm1.8-11h12.4l-1.05-3H6.85L5.8 13ZM5 15v5v-5Zm2.5 4q.625 0 1.063-.438T9 17.5q0-.625-.438-1.063T7.5 16q-.625 0-1.063.438T6 17.5q0 .625.438 1.063T7.5 19Zm9 0q.625 0 1.063-.438T18 17.5q0-.625-.438-1.063T16.5 16q-.625 0-1.063.438T15 17.5q0 .625.438 1.063T16.5 19ZM5 20h14v-5H5v5Z"/>`),
		g.Group(children),
	)
}