package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kungfu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSKungfu0"><g fill="none" stroke="#fff" stroke-miterlimit="2" stroke-width="4"><path fill="#fff" d="M11 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m20 18l8 6l-3 18m3-18l16-13"/><path stroke-linecap="round" stroke-linejoin="round" d="M28 16.8L27 8l-7 10l-4 9l-6 3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSKungfu0)"/>`),
		g.Group(children),
	)
}