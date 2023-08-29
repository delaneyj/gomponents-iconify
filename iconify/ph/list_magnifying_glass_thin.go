package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListMagnifyingGlassThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36 64a4 4 0 0 1 4-4h176a4 4 0 0 1 0 8H40a4 4 0 0 1-4-4Zm4 68h72a4 4 0 0 0 0-8H40a4 4 0 0 0 0 8Zm88 56H40a4 4 0 0 0 0 8h88a4 4 0 0 0 0-8Zm106.83 6.83a4 4 0 0 1-5.66 0l-22.72-22.72a36.06 36.06 0 1 1 5.66-5.66l22.72 22.72a4 4 0 0 1 0 5.66ZM184 172a28 28 0 1 0-28-28a28 28 0 0 0 28 28Z"/>`),
		g.Group(children),
	)
}