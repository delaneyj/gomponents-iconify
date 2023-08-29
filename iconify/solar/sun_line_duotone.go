package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="6"/><path stroke-linecap="round" d="M12 2v1m0 18v1m10-10h-1M3 12H2"/><path stroke-linecap="round" d="m19.07 4.93l-.392.393M5.322 18.678l-.393.393m14.141-.001l-.392-.393M5.322 5.322l-.393-.393" opacity=".5"/></g>`),
		g.Group(children),
	)
}