package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.175 1.4L7.6 0l8.575 8.575q.575.575.575 1.425t-.575 1.425l-4.75 4.75q-.575.575-1.425.575t-1.425-.575l-4.75-4.75Q3.25 10.85 3.25 10t.575-1.425L8.575 3.8l-2.4-2.4ZM10 5.225L5.225 10h9.55L10 5.225ZM19 17q-.825 0-1.413-.588T17 15q0-.525.313-1.125T18 12.75q.225-.3.475-.625T19 11.5q.275.3.525.625t.475.625q.375.525.688 1.125T21 15q0 .825-.588 1.413T19 17ZM2 24v-4h20v4H2Z"/>`),
		g.Group(children),
	)
}