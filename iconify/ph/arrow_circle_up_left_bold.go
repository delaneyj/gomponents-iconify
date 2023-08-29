package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUpLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm40.49-60.49a12 12 0 0 1-17 17L108 125v19a12 12 0 0 1-24 0V96a12 12 0 0 1 12-12h48a12 12 0 0 1 0 24h-19Z"/>`),
		g.Group(children),
	)
}