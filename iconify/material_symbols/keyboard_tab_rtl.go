package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardTabRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18H2V6h2v12Zm8 0l-6-6l6-6l1.4 1.4L9.825 11H22v2H9.825l3.6 3.6L12 18Z"/>`),
		g.Group(children),
	)
}