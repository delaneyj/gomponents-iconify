package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm64.12 133.14a8 8 0 0 1-8.24 13.72L136 142.13V194a8 8 0 0 1-16 0v-51.87l-47.88 28.73a8 8 0 0 1-8.24-13.72L112.45 128L63.88 98.86a8 8 0 0 1 8.24-13.72L120 113.87V62a8 8 0 0 1 16 0v51.87l47.88-28.73a8 8 0 1 1 8.24 13.72L143.55 128Z"/>`),
		g.Group(children),
	)
}