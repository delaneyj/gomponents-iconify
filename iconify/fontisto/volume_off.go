package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.883.133L7.09 5.188v13.625l8.793 5.055c.724.416 1.571-.202 1.571-1.145V1.279c0-.945-.847-1.564-1.571-1.145zM6 5.615l-4.926.01A1.184 1.184 0 0 0 .001 6.888v-.004V17.14a1.167 1.167 0 0 0 1.059 1.246H6V5.614zm23.196 2.893l-.887-.884l-3.491 3.491l-3.492-3.491l-.884.884l3.491 3.491l-3.491 3.492l.884.887l3.491-3.491l3.492 3.491l.887-.887L25.702 12z"/>`),
		g.Group(children),
	)
}