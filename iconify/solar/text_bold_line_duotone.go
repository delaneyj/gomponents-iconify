package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBoldLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5 4.609A2.609 2.609 0 0 1 7.609 2H12a5 5 0 0 1 0 10H5V4.609Z"/><path d="M5 12h9a5 5 0 0 1 0 10H7.059A2.059 2.059 0 0 1 5 19.941V12Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}