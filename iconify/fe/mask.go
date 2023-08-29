package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMask0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMask1" fill="currentColor"><path id="feMask2" d="M5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2Zm14 16c0-7.732-6.268-14-14-14v14h14Z"/></g></g>`),
		g.Group(children),
	)
}