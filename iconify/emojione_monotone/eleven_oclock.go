package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevenOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm2 36h-4v-2.555c-1.19-.693-2-1.969-2-3.445c0-.701.196-1.351.514-1.924L21 17.932L24.584 16L30 24.755V6h4v22.555A3.98 3.98 0 0 1 36 32c0 .637-.162 1.23-.427 1.764L37 36.07l-3 1.615V38z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}