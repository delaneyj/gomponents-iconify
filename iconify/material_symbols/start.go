package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Start(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V6h2v12H2Zm14 0l-1.425-1.4l3.6-3.6H6v-2h12.175L14.6 7.4L16 6l6 6l-6 6Z"/>`),
		g.Group(children),
	)
}