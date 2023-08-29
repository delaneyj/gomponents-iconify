package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentHorizontalBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignmentHorizontalBottom0"><g fill="none" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linejoin="round" rx="3"/><path stroke="#000" d="M22 36h4m-8-6h12m-10-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignmentHorizontalBottom0)"/>`),
		g.Group(children),
	)
}