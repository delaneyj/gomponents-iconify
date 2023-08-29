package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperFolding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNewspaperFolding0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m22 44l-1-8"/><path fill="#555" d="M42 44V12H26l1 8l1 8l1 8l-7 8h20Z"/><path d="M28 28h5m-6-8h6"/><path fill="#555" d="M6 4h19l1 8l1 8l1 8l1 8H6V4Z"/><path d="M12 12h7m-7 8h8m-8 8h9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNewspaperFolding0)"/>`),
		g.Group(children),
	)
}