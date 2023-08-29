package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoCoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.623 5.248A9.964 9.964 0 0 0 2 12c0 5.523 4.477 10 10 10a9.962 9.962 0 0 0 6.615-2.5m2.687-3.822A9.974 9.974 0 0 0 22 12c0-5.523-4.477-10-10-10c-1.231 0-2.41.223-3.5.63"/><path d="M9 15c.644.86 1.843 1.35 3 1.391c1.114.04 2.19-.336 2.697-1.198M12 16.391V18.5m-2.5-9c0 1.181.852 1.665 1.886 2M15 8.5c-.685-.685-1.891-1.161-3-1.191V5.5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}