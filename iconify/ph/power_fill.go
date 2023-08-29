package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104 104 0 0 0 128 24Zm-8 40a8 8 0 0 1 16 0v64a8 8 0 0 1-16 0Zm8 144A80 80 0 0 1 83.55 61.48a8 8 0 1 1 8.9 13.29a64 64 0 1 0 71.1 0a8 8 0 1 1 8.9-13.29A80 80 0 0 1 128 208Z"/>`),
		g.Group(children),
	)
}