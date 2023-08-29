package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 17h5l1.67-2.386m3.66-5.227L15 7h6"/><path d="m18 4l3 3l-3 3M3 7h5l7 10h6"/><path d="m18 20l3-3l-3-3"/></g>`),
		g.Group(children),
	)
}