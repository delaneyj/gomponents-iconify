package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIpad0"><g fill="none" stroke-width="4"><rect width="30" height="38" x="9" y="5" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M22 36h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIpad0)"/>`),
		g.Group(children),
	)
}