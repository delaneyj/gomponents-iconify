package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyleftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 56a72 72 0 1 0 72 72a72.08 72.08 0 0 0-72-72Zm0 120a47.66 47.66 0 0 1-38.4-19.19a8 8 0 0 1 12.8-9.61a32 32 0 1 0 0-38.4a8 8 0 0 1-12.8-9.61A48 48 0 1 1 128 176Zm0-152a104 104 0 1 0 104 104A104 104 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Z"/>`),
		g.Group(children),
	)
}