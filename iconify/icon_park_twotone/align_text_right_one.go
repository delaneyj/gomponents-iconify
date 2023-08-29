package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextRightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignTextRightOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" d="M34 24H22m12-9H14m20 18H16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignTextRightOne0)"/>`),
		g.Group(children),
	)
}