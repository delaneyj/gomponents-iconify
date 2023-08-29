package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLeftSquare0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" d="M29 34L19 24l10-10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLeftSquare0)"/>`),
		g.Group(children),
	)
}