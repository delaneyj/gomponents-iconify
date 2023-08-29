package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextCenterOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAlignTextCenterOne0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M30 24H18m16-9H14m20 18H14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAlignTextCenterOne0)"/>`),
		g.Group(children),
	)
}