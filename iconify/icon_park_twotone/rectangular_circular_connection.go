package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangularCircularConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRectangularCircularConnection0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M12 19a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm5 12H7v10h10V31Z"/><path stroke-linecap="round" d="M25.68 13H42v23H25"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRectangularCircularConnection0)"/>`),
		g.Group(children),
	)
}