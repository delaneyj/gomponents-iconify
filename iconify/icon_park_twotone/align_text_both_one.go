package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextBothOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAlignTextBothOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" d="M34 24H14m20-9H14m20 18H14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAlignTextBothOne0)"/>`),
		g.Group(children),
	)
}