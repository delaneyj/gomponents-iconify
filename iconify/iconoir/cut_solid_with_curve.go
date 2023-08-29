package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CutSolidWithCurve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.528 7.293L9 10.333M22 2h-2m-8 10v-2a8.004 8.004 0 0 1 5.5-7.602M22 12h-2m-8 10v-2a8.004 8.004 0 0 1 5.5-7.602"/><path d="m12 22l-8.691-4.828A.6.6 0 0 1 3 16.647V7.353a.6.6 0 0 1 .309-.524l8.4-4.667a.6.6 0 0 1 .582 0L15 3.667M12 12L3.528 7.293M12 21v-9m3 1.5V4"/></g>`),
		g.Group(children),
	)
}