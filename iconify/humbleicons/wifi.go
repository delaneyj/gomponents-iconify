package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8.437c-5.115-4.583-12.885-4.583-18 0m15.333 2.982a9.501 9.501 0 0 0-12.666 0m10 2.981a5.5 5.5 0 0 0-7.334 0M12 18.5l1-1.118a1.5 1.5 0 0 0-2 0l1 1.118Z"/>`),
		g.Group(children),
	)
}