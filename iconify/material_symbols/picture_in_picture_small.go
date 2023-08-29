package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-2h18V4h2v14q0 .825-.588 1.413T20 20H2Zm8-4v-6h8v6h-8Z"/>`),
		g.Group(children),
	)
}