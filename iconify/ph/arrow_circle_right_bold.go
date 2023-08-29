package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm48.49-92.49a12 12 0 0 1 0 17l-32 32a12 12 0 1 1-17-17L139 140H88a12 12 0 0 1 0-24h51l-11.52-11.51a12 12 0 1 1 17-17Z"/>`),
		g.Group(children),
	)
}