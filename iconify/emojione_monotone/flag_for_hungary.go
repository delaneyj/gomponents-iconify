package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForHungary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM5.787 22.166h52.426A27.87 27.87 0 0 1 60 32c0 3.459-.635 6.771-1.787 9.833H5.787A27.864 27.864 0 0 1 4 32c0-3.459.635-6.772 1.787-9.834z"/>`),
		g.Group(children),
	)
}