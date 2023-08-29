package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swatches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 180a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm152-23.81V208a16 16 0 0 1-16 16H68a46.36 46.36 0 0 1-7.94-.68a44 44 0 0 1-35.43-50.95l25-143.13a15.94 15.94 0 0 1 18.47-13L122.84 26a16 16 0 0 1 12.92 18.52l-12.08 69L191.49 89a16 16 0 0 1 20.45 9.52L231 150.69a18.35 18.35 0 0 1 1 5.5ZM95 184.87l25-143.13L65.46 32l-25 143.1a28 28 0 0 0 22.44 32.47A27.29 27.29 0 0 0 83.46 203A27.84 27.84 0 0 0 95 184.87ZM108.78 195L216 156.11L196.92 104l-76.42 27.7l-9.78 55.92a44.63 44.63 0 0 1-1.94 7.38ZM216 173.12L119.74 208H216Z"/>`),
		g.Group(children),
	)
}