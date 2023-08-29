package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PensiveFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" d="M36 13c-12.682 0-23 10.318-23 23s10.318 23 23 23s23-10.318 23-23s-10.318-23-23-23z"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="23"/><path stroke-linecap="round" stroke-linejoin="round" d="M27 46h18M23 33.077c.413.344 2.246 1.792 4.91 1.637c2.16-.126 3.61-1.233 4.09-1.637m9 0c.413.344 2.246 1.792 4.91 1.637c2.16-.126 3.61-1.233 4.09-1.637M19.939 27.25c.487.228 2.627 1.16 5.163.333c2.058-.672 3.178-2.112 3.54-2.624m24 2.291c-.487.228-2.628 1.16-5.164.333c-2.057-.672-3.177-2.112-3.54-2.624"/></g>`),
		g.Group(children),
	)
}