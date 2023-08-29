package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sanitizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 8q-.625 0-1.063-.438T15.5 6.5q0-.425.425-1.125T17 4q.65.675 1.075 1.375T18.5 6.5q0 .625-.438 1.063T17 8Zm2.5 7q-1.05 0-1.775-.725T17 12.5q0-.875.775-2.163T19.5 8q.95 1.05 1.725 2.337T22 12.5q0 1.05-.725 1.775T19.5 15ZM9 18h2v-2h2v-2h-2v-2H9v2H7v2h2v2Zm-3 4q-.825 0-1.413-.588T4 20v-8q0-2.25 1.425-3.9T9 6.1V4H7V2h6q.85 0 1.6.263T16 3l-1.45 1.45q-.35-.2-.738-.325T13 4h-2v2.1q2.15.35 3.575 2T16 12v8q0 .825-.588 1.413T14 22H6Z"/>`),
		g.Group(children),
	)
}