package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentLeftTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignmentLeftTop0"><g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="3"/><path stroke="#000" d="M12 24h4m-4-6h12m-12-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignmentLeftTop0)"/>`),
		g.Group(children),
	)
}