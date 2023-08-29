package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationGraphLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M2 2h20M9 10.5l1.293-1.293c.333-.333.5-.5.707-.5c.207 0 .374.167.707.5l.586.586c.333.333.5.5.707.5c.207 0 .374-.167.707-.5L15 8.5M12 21v-4m-2 5l2-1m2 1l-2-1"/><path d="M20 2v8.5c0 3.064 0 4.596-1.004 5.548s-2.62.952-5.853.952h-2.286c-3.232 0-4.849 0-5.853-.952C4 15.096 4 13.564 4 10.5V2"/></g>`),
		g.Group(children),
	)
}