package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentOutlook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2.998 9V1H17.5L21 4.5V23H2M16 1v5h5M7.75 15.75C7.75 13.5 6.5 12 5 12s-2.75 1.5-2.75 3.75S3.5 19.5 5 19.5s2.75-1.5 2.75-3.75ZM5 12c2.425 0 3 2.5 3 3.75s-.5 3.75-3 3.75s-3-2.5-3-3.75S2.559 12 5 12Z"/>`),
		g.Group(children),
	)
}