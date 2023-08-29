package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M112 116a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm44-12a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm68 16v96a8 8 0 0 1-13.07 6.19l-24.26-19.85l-24.27 19.85a8 8 0 0 1-10.13 0L128 202.34l-24.27 19.85a8 8 0 0 1-10.13 0l-24.27-19.85l-24.26 19.85A8 8 0 0 1 32 216v-96a96 96 0 0 1 192 0Zm-16 0a80 80 0 0 0-160 0v79.12l16.27-13.31a8 8 0 0 1 10.13 0l24.27 19.85l24.26-19.85a8 8 0 0 1 10.14 0l24.26 19.85l24.27-19.85a8 8 0 0 1 10.13 0L208 199.12Z"/>`),
		g.Group(children),
	)
}