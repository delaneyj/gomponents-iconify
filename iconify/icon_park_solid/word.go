package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Word(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWord0"><g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m14 16l4 16l6-13l6 13l4-16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWord0)"/>`),
		g.Group(children),
	)
}