package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowRightDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m224 160l-48 48l-48-48Z" opacity=".2"/><path d="M231.39 156.94A8 8 0 0 0 224 152h-40V64a8 8 0 0 0-8-8H32a8 8 0 0 0 0 16h136v80h-40a8 8 0 0 0-5.66 13.66l48 48a8 8 0 0 0 11.32 0l48-48a8 8 0 0 0 1.73-8.72ZM176 196.69L147.31 168h57.38Z"/></g>`),
		g.Group(children),
	)
}