package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Margin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMargin0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" d="M34 6v36M14 6v36m17 0h6m-26 0h6M11 6h6m13 0h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMargin0)"/>`),
		g.Group(children),
	)
}