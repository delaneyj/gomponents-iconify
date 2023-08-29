package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DsCam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="33.331" height="26.345" x="4.5" y="10.828" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.28"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.04 30.304l3.648 2.106a1.208 1.208 0 0 0 1.812-1.046V16.636a1.208 1.208 0 0 0-1.812-1.046l-3.856 2.226"/><rect width="21.03" height="5.319" x="10.651" y="17.34" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.66"/>`),
		g.Group(children),
	)
}