package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapArrowRightLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M5.527 12c0-.25-.061-.5-.184-.73L2.191 5.368c-.74-1.388.81-2.87 2.306-2.202l16.51 7.362c.662.296.993.884.993 1.473"/><path d="M5.527 12c0 .25-.061.5-.184.73l-3.152 5.903c-.74 1.388.81 2.87 2.306 2.202l16.51-7.363C21.67 13.178 22 12.59 22 12" opacity=".5"/></g>`),
		g.Group(children),
	)
}