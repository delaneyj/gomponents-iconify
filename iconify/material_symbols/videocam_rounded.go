package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l3.15-3.15q.225-.225.537-.113T22 7.7v8.6q0 .35-.313.463t-.537-.113L18 13.5V18q0 .825-.588 1.413T16 20H4Z"/>`),
		g.Group(children),
	)
}