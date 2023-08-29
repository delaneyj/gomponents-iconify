package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudLightningBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 12a80.22 80.22 0 0 0-73.61 48.36A56.76 56.76 0 0 0 76 60a56 56 0 0 0 0 112h30.81l-13.1 21.82A12 12 0 0 0 104 212h18.81l-13.1 21.82a12 12 0 1 0 20.58 12.35l24-40A12 12 0 0 0 144 188h-18.81l9.6-16H156a80 80 0 0 0 0-160Zm0 136H76a32 32 0 0 1 0-64h.28c-.11 1.1-.2 2.2-.26 3.3a12 12 0 1 0 24 1.39A56.06 56.06 0 1 1 156 148Z"/>`),
		g.Group(children),
	)
}