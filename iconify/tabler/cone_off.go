package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.396 16.384L12.87 2.507a1 1 0 0 0-1.74 0L9.504 5.505M8.097 8.099L3 17.497v.5C3 19.657 7.03 21 12 21c3.202 0 6.014-.558 7.609-1.398M3 3l18 18"/>`),
		g.Group(children),
	)
}