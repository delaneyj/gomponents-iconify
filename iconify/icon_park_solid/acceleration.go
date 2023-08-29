package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acceleration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAcceleration0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path d="M42 35h-8m-3 7h-4"/><path fill="#fff" stroke-linejoin="round" d="m14 22l-6-5H4s5.486 10 8 10h32l-6-5H14Z"/><path stroke-linejoin="round" d="m30 22l-11.34-8H16l3 8m11 5L17.2 39H14l4.267-12"/><path fill="#fff" stroke-linejoin="round" d="M32 11c0 1-3 2-3 2h10s2.886 0 3.745-2.286C43.63 8.36 42.045 5 39.022 5C36 5 36 8 36 8s-1.855-.571-3 0s-1 2-1 3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAcceleration0)"/>`),
		g.Group(children),
	)
}