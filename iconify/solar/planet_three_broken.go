package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlanetThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M3 8.008S5.937 11 10.437 11M21 8.008s-1.917-.138-4.5.748c-.797.273-1.303.695-2.5 1.244M3 14.008s.768-.05 2 .107m2.903.64C9.401 15.228 11.172 17 14.51 17c3.012 0 5.381-1.129 6.761-2"/><path d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}