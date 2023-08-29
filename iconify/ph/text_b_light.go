package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M166.69 116.41A42 42 0 0 0 140 42H72a6 6 0 0 0-6 6v152a6 6 0 0 0 6 6h80a46 46 0 0 0 14.69-89.59ZM78 54h62a30 30 0 0 1 0 60H78Zm74 140H78v-68h74a34 34 0 0 1 0 68Z"/>`),
		g.Group(children),
	)
}