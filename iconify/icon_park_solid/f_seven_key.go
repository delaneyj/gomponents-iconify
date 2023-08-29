package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FSevenKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFSevenKey0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" d="M26 16h8l-6 16m-7-16h-7v16m0-8h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFSevenKey0)"/>`),
		g.Group(children),
	)
}