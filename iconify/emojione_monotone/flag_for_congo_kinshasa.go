package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCongoKinshasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m14.585 22.166l3.647-2.769l3.649 2.769l-1.377-4.502l3.629-2.874h-4.498l-1.403-4.426l-1.362 4.426h-4.538l3.63 2.874z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m12.502 4.957L6.957 44.502a27.724 27.724 0 0 1-.858-1.879L42.624 6.1c.639.262 1.265.549 1.878.857M19.5 57.045L57.045 19.5c.307.614.594 1.24.857 1.878L21.379 57.901a27.704 27.704 0 0 1-1.879-.856M32 4c2.982 0 5.855.473 8.553 1.341L5.341 40.553A27.878 27.878 0 0 1 4 32C4 16.561 16.561 4 32 4m0 56c-2.981 0-5.854-.473-8.551-1.34L58.66 23.449A27.895 27.895 0 0 1 60 32c0 15.439-12.561 28-28 28"/>`),
		g.Group(children),
	)
}