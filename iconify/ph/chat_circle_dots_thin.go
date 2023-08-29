package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCircleDotsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M136 128a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-52-8a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm88 0a8 8 0 1 0 8 8a8 8 0 0 0-8-8Zm56 8a100 100 0 0 1-148.5 87.47l-35.69 11.9a12 12 0 0 1-15.18-15.18l11.9-35.69A100 100 0 1 1 228 128Zm-8 0a92 92 0 1 0-171.65 46.07a4 4 0 0 1 .33 3.27l-12.46 37.38a4 4 0 0 0 5.06 5.06l37.38-12.46a3.93 3.93 0 0 1 1.27-.21a4.05 4.05 0 0 1 2 .54A92 92 0 0 0 220 128Z"/>`),
		g.Group(children),
	)
}