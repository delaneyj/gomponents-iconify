package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinusLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536M3.465 20.536C2 19.07 2 16.713 2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464M3.465 20.536l17.07-17.072"/><path stroke-linecap="round" d="M18 17h-5m-2.5-9H8m0 0H5.5M8 8V5.5M8 8v2.5" opacity=".5"/></g>`),
		g.Group(children),
	)
}