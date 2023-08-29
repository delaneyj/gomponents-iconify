package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReminderHandGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m17.5 12l2.004 2.672a2 2 0 0 1-.126 2.552l-3.784 4.127a1.998 1.998 0 0 1-1.473.649H9.5c-2.358 0-3.622-2.575-3.982-3.93a.55.55 0 0 1-.018-.143V9.43c0-2.286 3-2.286 3 0V10"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.5 10V8.286c0-2.286-3-2.286-3 0V10m6 0V7.5c0-2.286-3-2.286-3 0c0 0 0 0 0 0V10m3 0V3.499A1.5 1.5 0 0 1 16 2v0a1.5 1.5 0 0 1 1.5 1.5V15"/><path d="M17.563 6.5h2.062C20.5 6.5 21 6.078 21 5.25C21 4.422 20.5 4 19.625 4H12.25C11.56 4 11 4.56 11 5.25v.25a1 1 0 0 0 1 1"/></g>`),
		g.Group(children),
	)
}