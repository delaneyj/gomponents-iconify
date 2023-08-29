package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepingCircleBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M6.5 11c.567.63 1.256 1 2 1s1.433-.37 2-1m3 0c.567.63 1.256 1 2 1s1.433-.37 2-1"/><path fill="currentColor" d="M13 16a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 4l3.464-2L19 7.464l3.464-2m-8.416.036l1.732 1l-2.732.732l1.732 1"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 12c0 5.523-4.477 10-10 10a9.955 9.955 0 0 1-5-1.338M12 2C6.477 2 2 6.477 2 12c0 1.821.487 3.53 1.338 5"/></g>`),
		g.Group(children),
	)
}