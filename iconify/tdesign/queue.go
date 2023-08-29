package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Queue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 2.001l.003 18.418L20 18.415V4.001l-12.999.001v-2L22 2.001Zm-20 4h16v16H2V6Zm2 2v12h12V8H4Zm7 1.5V13h3.5v2H11v3.5H9V15H5.5v-2H9V9.5h2Z"/>`),
		g.Group(children),
	)
}