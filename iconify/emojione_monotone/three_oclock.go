package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m17.998 32H35.445A3.98 3.98 0 0 1 34 35.445V38h-4v-2.555A3.996 3.996 0 0 1 28.555 34h-2.553L26 30h2.555c.35-.6.845-1.096 1.445-1.445V6h4v22.555A3.98 3.98 0 0 1 35.445 30H50l-.002 4"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}