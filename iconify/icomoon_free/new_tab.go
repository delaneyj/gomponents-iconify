package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 1v12h12V1H3zm11 11H4V2h10v10zM2 14V3.5l-1-1V15h12.5l-1-1H2z"/><path fill="currentColor" d="M5.5 4L8 6.5l-3 3L6.5 11l3-3l2.5 2.5V4z"/>`),
		g.Group(children),
	)
}