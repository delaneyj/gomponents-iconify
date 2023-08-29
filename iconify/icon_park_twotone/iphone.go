package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIphone0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="26" height="40" x="11" y="4" fill="#555" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 10h4m-6 28h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIphone0)"/>`),
		g.Group(children),
	)
}