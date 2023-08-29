package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBreakFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M239.81 99.5c-5.19 67.42-103.7 121.23-108 123.54a8 8 0 0 1-7.58 0C119.8 220.67 16 164 16 94a62 62 0 0 1 96.47-51.55a4 4 0 0 1 .61 6.17L99.72 62a8 8 0 0 0 0 11.31l32.53 32.53L111 127a8 8 0 1 0 11.31 11.31l26.88-26.87a8 8 0 0 0 0-11.31l-32.49-32.5l17.47-17.47A61.63 61.63 0 0 1 178.41 32c36.32.23 64.18 31.29 61.4 67.5Z"/>`),
		g.Group(children),
	)
}