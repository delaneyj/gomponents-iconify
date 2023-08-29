package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperFolding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m22 44l-1-8m21 8V12H26l1 8l1 8l1 8l-7 8h20ZM28 28h5m-6-8h6"/><path d="M6 4h19l1 8l1 8l1 8l1 8H6V4Zm6 8h7m-7 8h8m-8 8h9"/></g>`),
		g.Group(children),
	)
}