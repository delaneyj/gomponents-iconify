package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHandleUp0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linecap="round" rx="3"/><path fill="#000" stroke="#000" d="m14 28l10-10l10 10H14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHandleUp0)"/>`),
		g.Group(children),
	)
}