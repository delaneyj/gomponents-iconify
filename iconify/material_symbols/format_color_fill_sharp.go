package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorFillSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17.6L2.4 10l6.175-6.2l-2.4-2.4L7.6 0l10 10l-7.6 7.6Zm0-12.375L5.225 10h9.55L10 5.225ZM19 17q-.825 0-1.413-.588T17 15q0-.525.313-1.125T18 12.75q.225-.3.475-.625T19 11.5q.275.3.525.625t.475.625q.375.525.688 1.125T21 15q0 .825-.588 1.413T19 17ZM2 24v-4h20v4H2Z"/>`),
		g.Group(children),
	)
}