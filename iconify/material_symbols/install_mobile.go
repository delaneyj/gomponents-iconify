package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h7v5H7v12h10v-2h2v5q0 .825-.588 1.413T17 23H7Zm11-9l-5-5l1.4-1.4l2.6 2.6V3h2v7.2l2.6-2.6L23 9l-5 5Z"/>`),
		g.Group(children),
	)
}