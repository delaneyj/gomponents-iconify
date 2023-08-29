package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallDesktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h8v2H4v12h16v-3h2v3q0 .825-.588 1.413T20 19h-4v2H8Zm9-7l-5-5l1.4-1.4l2.6 2.575V3h2v7.175L20.6 7.6L22 9l-5 5Z"/>`),
		g.Group(children),
	)
}