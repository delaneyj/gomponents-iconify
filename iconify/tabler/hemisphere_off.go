package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HemisphereOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.588 6.603C4.41 7.15 3 8.02 3 9c0 1.657 4.03 3 9 3m3.72-.267C18.834 11.26 21 10.215 21 9c0-1.657-4.03-3-9-3c-.662 0-1.308.024-1.93.07"/><path d="M3 9a9 9 0 0 0 13.677 7.69m2.165-1.843A8.965 8.965 0 0 0 21 9M3 3l18 18"/></g>`),
		g.Group(children),
	)
}