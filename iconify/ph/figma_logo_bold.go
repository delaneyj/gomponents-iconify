package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M190.15 96A44 44 0 0 0 160 20H88a44 44 0 0 0-30.15 76a43.9 43.9 0 0 0 1.23 65.12A48 48 0 1 0 140 196v-28.83A44 44 0 0 0 190.15 96ZM180 64a20 20 0 0 1-20 20h-20V44h20a20 20 0 0 1 20 20ZM68 64a20 20 0 0 1 20-20h28v40H88a20 20 0 0 1-20-20Zm20 84a20 20 0 0 1 0-40h28v40H88Zm28 48a24 24 0 1 1-24-24h24Zm44-48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}