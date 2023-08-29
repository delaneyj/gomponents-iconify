package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorFillRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16.75q-.4 0-.763-.15t-.662-.425l-4.75-4.75q-.275-.3-.425-.663T3.25 10q0-.4.15-.775t.425-.65L8.575 3.8l-1.7-1.7Q6.6 1.825 6.6 1.413T6.9.7q.3-.275.7-.287T8.3.7l7.875 7.875q.3.275.438.65t.137.775q0 .4-.138.763t-.437.662l-4.75 4.75q-.275.275-.65.425t-.775.15Zm0-11.525L5.225 10h9.55L10 5.225ZM19 17q-.825 0-1.413-.588T17 15q0-.525.313-1.125T18 12.75q.225-.3.475-.625T19 11.5q.275.3.525.625t.475.625q.375.525.688 1.125T21 15q0 .825-.588 1.413T19 17ZM4 24q-.825 0-1.413-.588T2 22q0-.825.588-1.413T4 20h16q.825 0 1.413.588T22 22q0 .825-.588 1.413T20 24H4Z"/>`),
		g.Group(children),
	)
}