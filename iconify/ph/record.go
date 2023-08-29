package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm0-160a72 72 0 1 0 72 72a72.08 72.08 0 0 0-72-72Zm0 128a56 56 0 1 1 56-56a56.06 56.06 0 0 1-56 56Z"/>`),
		g.Group(children),
	)
}