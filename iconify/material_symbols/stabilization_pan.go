package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StabilizationPan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 17l-1.4-1.4l2.575-2.6H6v-2h8.175L11.6 8.4L13 7l5 5Zm3 4v-2h3v-3h2v5ZM3 21v-5h2v3h3v2ZM3 8V3h5v2H5v3Zm16 0V5h-3V3h5v5Z"/>`),
		g.Group(children),
	)
}