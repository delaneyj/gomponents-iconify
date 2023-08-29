package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm8 191.63V184a8 8 0 0 0-16 0v31.63A88.13 88.13 0 0 1 40.37 136H72a8 8 0 0 0 0-16H40.37A88.13 88.13 0 0 1 120 40.37V72a8 8 0 0 0 16 0V40.37A88.13 88.13 0 0 1 215.63 120H184a8 8 0 0 0 0 16h31.63A88.13 88.13 0 0 1 136 215.63Z"/>`),
		g.Group(children),
	)
}