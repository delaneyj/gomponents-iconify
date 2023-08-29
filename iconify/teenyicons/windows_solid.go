package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14.814.111A.5.5 0 0 1 15 .5V7H7V1.596L14.395.01a.5.5 0 0 1 .42.1ZM6 1.81L.395 3.011A.5.5 0 0 0 0 3.5V7h6V1.81ZM0 8v4.5a.5.5 0 0 0 .43.495l5.57.796V8H0Zm7 5.934l7.43 1.061A.5.5 0 0 0 15 14.5V8H7v5.934Z"/>`),
		g.Group(children),
	)
}