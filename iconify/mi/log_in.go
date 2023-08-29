package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14 19a1 1 0 1 0 0 2h5a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-5a1 1 0 1 0 0 2h5v14h-5z"/><path d="M15.714 12.7a.996.996 0 0 0 .286-.697v-.006a.997.997 0 0 0-.293-.704l-4-4a1 1 0 1 0-1.414 1.414L12.586 11H3a1 1 0 1 0 0 2h9.586l-2.293 2.293a1 1 0 1 0 1.414 1.414l4-4l.007-.007z"/></g>`),
		g.Group(children),
	)
}