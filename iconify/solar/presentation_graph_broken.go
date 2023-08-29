package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationGraphBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M2 2h20M9 10.5l1.293-1.293c.333-.333.5-.5.707-.5c.207 0 .374.167.707.5l.586.586c.333.333.5.5.707.5c.207 0 .374-.167.707-.5L15 8.5M12 21v-4m-2 5l2-1m2 1l-2-1"/><path d="M4 2v8.5c0 3.064 0 4.596 1.004 5.548s2.62.952 5.853.952h2.286c3.232 0 4.849 0 5.853-.952C20 15.096 20 13.564 20 10.5V10m0-8v4"/></g>`),
		g.Group(children),
	)
}