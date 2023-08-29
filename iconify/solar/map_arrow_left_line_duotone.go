package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapArrowLeftLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M18.473 12c0 .25.061.5.184.73l3.152 5.903c.74 1.388-.81 2.87-2.306 2.202l-16.51-7.363C2.33 13.178 2 12.59 2 12"/><path d="M18.473 12c0-.25.061-.5.184-.73l3.152-5.903c.74-1.388-.81-2.87-2.306-2.202l-16.51 7.362C2.33 10.823 2 11.412 2 12" opacity=".5"/></g>`),
		g.Group(children),
	)
}