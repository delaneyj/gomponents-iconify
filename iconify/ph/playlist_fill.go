package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M32 64a8 8 0 0 1 8-8h176a8 8 0 0 1 0 16H40a8 8 0 0 1-8-8Zm8 72h120a8 8 0 0 0 0-16H40a8 8 0 0 0 0 16Zm72 48H40a8 8 0 0 0 0 16h72a8 8 0 0 0 0-16Zm130.3-67.66l-40-12A8 8 0 0 0 192 112v52.31A32 32 0 1 0 208 192v-69.25l29.7 8.91a8 8 0 1 0 4.6-15.32Z"/>`),
		g.Group(children),
	)
}