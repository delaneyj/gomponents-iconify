package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowLeftDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m128 160l-48 48l-48-48Z" opacity=".2"/><path d="M224 56H80a8 8 0 0 0-8 8v88H32a8 8 0 0 0-5.66 13.66l48 48a8 8 0 0 0 11.32 0l48-48A8 8 0 0 0 128 152H88V72h136a8 8 0 0 0 0-16ZM80 196.69L51.31 168h57.38Z"/></g>`),
		g.Group(children),
	)
}