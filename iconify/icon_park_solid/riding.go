package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Riding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRiding0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#fff" stroke-miterlimit="2" d="M33 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" d="M40 23h-7.63c-.56 0-1.1-.24-1.48-.65l-5.34-5.78a2.01 2.01 0 0 0-2.53-.36l-7.45 4.58a1 1 0 0 0 0 1.7l7.46 4.57c.59.36.96 1.01.96 1.71L24 38"/><circle cx="36.5" cy="36.5" r="7.5" fill="#fff"/><circle cx="11.5" cy="36.5" r="7.5" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRiding0)"/>`),
		g.Group(children),
	)
}