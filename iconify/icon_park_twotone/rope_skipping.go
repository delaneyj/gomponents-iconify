package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RopeSkipping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRopeSkipping0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 32V11a7 7 0 1 1 14 0v26a7 7 0 1 0 14 0V16"/><path fill="#555" d="M41 4v12h-6V4h6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 16h3m9 0h-3m0 0V4h-6v12m6 0h-6"/><path fill="#555" d="M7 44V32h6v12H7Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 32h-3m-9 0h3m0 0v12h6V32m-6 0h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRopeSkipping0)"/>`),
		g.Group(children),
	)
}