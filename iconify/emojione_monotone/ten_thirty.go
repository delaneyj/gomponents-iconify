package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm3.143 36L34 36.857V58h-4V35.445c-1.19-.693-2-1.969-2-3.445c0-.348.059-.68.142-1.001L18 20.857L20.857 18L30 27.143V26h4v2.554c1.19.692 2 1.968 2 3.446c0 .348-.058.68-.142 1.001L38 35.143L35.143 38z"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}