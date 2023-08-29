package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsCounterClockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 104H40a8 8 0 0 1-8-8V48a8 8 0 0 1 16 0v23.85C63.47 54.4 90.61 32 128 32c42.82 0 68.58 25.27 69.66 26.34a8 8 0 0 1-11.3 11.34C186.09 69.41 163.92 48 128 48c-35 0-60.23 24.65-72.56 40H88a8 8 0 0 1 0 16Zm128 48h-48a8 8 0 0 0 0 16h32.56C188.23 183.35 163 208 128 208c-35.92 0-58.09-21.41-58.36-21.68a8 8 0 0 0-11.3 11.34C59.42 198.73 85.18 224 128 224c37.39 0 64.53-22.4 80-39.85V208a8 8 0 0 0 16 0v-48a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}