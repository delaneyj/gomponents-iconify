package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36 64a4 4 0 0 1 4-4h176a4 4 0 0 1 0 8H40a4 4 0 0 1-4-4Zm4 68h120a4 4 0 0 0 0-8H40a4 4 0 0 0 0 8Zm72 56H40a4 4 0 0 0 0 8h72a4 4 0 0 0 0-8Zm131.83-62.85a4 4 0 0 1-5 2.68L204 117.38V192a28 28 0 1 1-8-19.57V112a4 4 0 0 1 5.15-3.83l40 12a4 4 0 0 1 2.68 4.98ZM196 192a20 20 0 1 0-20 20a20 20 0 0 0 20-20Z"/>`),
		g.Group(children),
	)
}