package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrenchFries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFrenchFries0"><g fill="none" stroke-width="4"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M21 22V12a2 2 0 0 0-2-2h-3a2 2 0 0 0-2 2v9m21 0v-7a2 2 0 0 0-2-2h-3a2 2 0 0 0-2 2v8m0 0V6a2 2 0 0 0-2-2h-3a2 2 0 0 0-2 2v16"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M7 18s7 5 17 5s17-5 17-5l-4.818 26H11.818L7 18Z"/><ellipse cx="24" cy="33" fill="#000" stroke="#000" rx="6" ry="3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFrenchFries0)"/>`),
		g.Group(children),
	)
}