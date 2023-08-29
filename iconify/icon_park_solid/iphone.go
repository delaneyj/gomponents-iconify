package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIphone0"><g fill="none" stroke-width="4"><rect width="26" height="40" x="11" y="4" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M22 10h4m-6 28h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIphone0)"/>`),
		g.Group(children),
	)
}