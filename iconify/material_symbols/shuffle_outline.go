package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20v-2h2.6l-3.175-3.175L14.85 13.4L18 16.55V14h2v6h-6Zm-8.6 0L4 18.6L16.6 6H14V4h6v6h-2V7.4L5.4 20Zm3.775-9.425L4 5.4L5.4 4l5.175 5.175l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}