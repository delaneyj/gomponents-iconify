package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vita(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="36.465" cy="15.816" r="6.035" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.115 38.388c-1.832 0-3.64-.83-4.827-2.406L6.713 19.277a6.036 6.036 0 0 1 9.644-7.259l12.575 16.706a6.036 6.036 0 0 1-4.817 9.663Z"/>`),
		g.Group(children),
	)
}