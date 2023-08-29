package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIpad0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="30" height="38" x="9" y="5" fill="#555" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 36h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIpad0)"/>`),
		g.Group(children),
	)
}