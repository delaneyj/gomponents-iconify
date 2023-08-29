package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shaving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShaving0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="20" height="27" x="14" y="17" fill="#fff" rx="2"/><path d="M19 12h10v5H19zm0 0V9c0-1 0-5 8-5h9v5h-7v3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShaving0)"/>`),
		g.Group(children),
	)
}