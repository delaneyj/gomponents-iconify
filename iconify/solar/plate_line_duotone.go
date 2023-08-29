package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlateLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15 4.001c3.114.01 4.765.108 5.828 1.17C22 6.344 22 8.23 22 12s0 5.657-1.172 6.828C19.657 20 17.771 20 14 20h-4c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12c0-3.771 0-5.657 1.172-6.828C4.235 4.109 5.886 4.01 9 4" opacity=".5"/><path d="M12 5V3m-4 7.5h8M8 14h5.5"/></g>`),
		g.Group(children),
	)
}