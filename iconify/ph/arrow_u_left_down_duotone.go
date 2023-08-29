package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m136 176l-48 48l-48-48Z" opacity=".2"/><path d="M144 24a64.07 64.07 0 0 0-64 64v80H40a8 8 0 0 0-5.66 13.66l48 48a8 8 0 0 0 11.32 0l48-48A8 8 0 0 0 136 168H96V88a48 48 0 0 1 96 0v88a8 8 0 0 0 16 0V88a64.07 64.07 0 0 0-64-64ZM88 212.69L59.31 184h57.38Z"/></g>`),
		g.Group(children),
	)
}