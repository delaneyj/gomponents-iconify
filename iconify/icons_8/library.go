package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 3.906l-.375.156l-12 5L3 9.345V12h2v11H3v5h26v-5h-2V12h2V9.344l-.625-.28l-12-5.002L16 3.906zm0 2.188L25.375 10H6.625L16 6.094zM7 12h2v11H7V12zm4 0h2v11h-2V12zm4 0h2v11h-2V12zm4 0h2v11h-2V12zm4 0h2v11h-2V12zM5 25h22v1H5v-1z"/>`),
		g.Group(children),
	)
}