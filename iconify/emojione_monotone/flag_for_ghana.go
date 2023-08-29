package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGhana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m22.167 29.542l6.05 4.788l-2.295 7.504L32 37.221l6.078 4.613l-2.295-7.504l6.051-4.788h-7.496L32 22.167l-2.27 7.375z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zM5.787 22.167h52.426A27.86 27.86 0 0 1 60 32c0 3.459-.635 6.771-1.787 9.834H5.787A27.873 27.873 0 0 1 4 32c0-3.459.635-6.772 1.787-9.833z"/>`),
		g.Group(children),
	)
}