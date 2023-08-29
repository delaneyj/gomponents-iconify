package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm46.21-143.58l-64 32a4.08 4.08 0 0 0-1.79 1.79l-32 64A4 4 0 0 0 80 180a4.05 4.05 0 0 0 1.79-.42l64-32a4.08 4.08 0 0 0 1.79-1.79l32-64a4 4 0 0 0-5.37-5.37ZM141 141l-52.08 26L115 115l52.08-26Z"/>`),
		g.Group(children),
	)
}