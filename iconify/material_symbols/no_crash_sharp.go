package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoCrashSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6.35l-2.825-2.8l1.4-1.425L12 3.55L15.55 0l1.4 1.4L12 6.35ZM3 24v-9l2.45-7h13.1L21 15v9h-3v-2H6v2H3Zm2.8-11h12.4l-1.05-3H6.85L5.8 13Zm1.7 6q.625 0 1.063-.438T9 17.5q0-.625-.438-1.063T7.5 16q-.625 0-1.063.438T6 17.5q0 .625.438 1.063T7.5 19Zm9 0q.625 0 1.063-.438T18 17.5q0-.625-.438-1.063T16.5 16q-.625 0-1.063.438T15 17.5q0 .625.438 1.063T16.5 19Z"/>`),
		g.Group(children),
	)
}