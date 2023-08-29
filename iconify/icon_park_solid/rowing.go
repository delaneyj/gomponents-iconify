package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rowing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRowing0"><g fill="none" stroke="#fff" stroke-miterlimit="2" stroke-width="4"><path fill="#fff" d="M30.02 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m42 31l-7 13m-2.99-4L26 29L9 40l12-20l23 5M4 40h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRowing0)"/>`),
		g.Group(children),
	)
}