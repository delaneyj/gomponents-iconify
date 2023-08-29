package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UfoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.95 9.01c3.02.739 5.05 2.123 5.05 3.714c0 1.08-.931 2.063-2.468 2.814m-3 1c-1.36.295-2.9.462-4.531.462c-5.52 0-10-1.909-10-4.276c0-1.59 2.04-2.985 5.07-3.724"/><path d="M14.69 10.686C16.078 10.331 17 9.71 17 9v-.035C17 6.223 14.761 4 12 4c-1.125 0-2.164.37-3 .992M7.293 7.289A4.925 4.925 0 0 0 7 8.965V9c0 .961 1.696 1.764 3.956 1.956M15 17l2 3m-8.5-3L7 20m5-6h.01M7 13h.01M17 13h.01M3 3l18 18"/></g>`),
		g.Group(children),
	)
}