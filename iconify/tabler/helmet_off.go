package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelmetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.633 4.654a9 9 0 0 1 11.718 11.7m-1.503 2.486A9.008 9.008 0 0 1 17.656 20H6.344a9 9 0 0 1-.185-13.847"/><path d="M20 9h-7m-2.768 1.246c.507 2 1.596 3.418 3.268 4.254c.524.262 1.07.49 1.64.683M3 3l18 18"/></g>`),
		g.Group(children),
	)
}