package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDRotate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3a7 7 0 0 1 7 7v4l-3-3m6 0l-3 3M8 15.5l-5-3l5-3l5 3V18l-5 3z"/><path d="M3 12.5V18l5 3m0-5.455l5-3.03"/></g>`),
		g.Group(children),
	)
}