package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignmentRightCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignmentRightCenter0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" stroke-linejoin="round" rx="3"/><path d="M32 30h4m-12-6h12m-8-6h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignmentRightCenter0)"/>`),
		g.Group(children),
	)
}