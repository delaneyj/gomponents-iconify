package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 64v120a16 16 0 0 1-16 16H32a16 16 0 0 0 16-16V64a8 8 0 0 1 8-8h160a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M88 112a8 8 0 0 1 8-8h80a8 8 0 0 1 0 16H96a8 8 0 0 1-8-8Zm8 40h80a8 8 0 0 0 0-16H96a8 8 0 0 0 0 16Zm136-88v120a24 24 0 0 1-24 24H32a24 24 0 0 1-24-23.89V88a8 8 0 0 1 16 0v96a8 8 0 0 0 16 0V64a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Zm-16 0H56v120a23.84 23.84 0 0 1-1.37 8H208a8 8 0 0 0 8-8Z"/></g>`),
		g.Group(children),
	)
}