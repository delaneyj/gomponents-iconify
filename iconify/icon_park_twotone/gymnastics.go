package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gymnastics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGymnastics0"><g fill="none" stroke="#fff" stroke-miterlimit="2" stroke-width="4"><path fill="#555" d="M24 22a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m23 29l-2 7l-9-2l-5 10m14-8l2 8h11M7 23l16 6l12-2l6.04-4.97M12 4c14-2 24 2 32 11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGymnastics0)"/>`),
		g.Group(children),
	)
}