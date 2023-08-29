package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm0-152a64 64 0 1 0 64 64a64.07 64.07 0 0 0-64-64Zm0 112a48 48 0 1 1 48-48a48.05 48.05 0 0 1-48 48Z"/>`),
		g.Group(children),
	)
}