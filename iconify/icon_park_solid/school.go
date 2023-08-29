package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func School(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSchool0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4 33a2 2 0 0 1 2-2h6v-7l12-8l12 8v7h6a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H4V33Z"/><path stroke="#fff" stroke-linecap="round" d="M24 6v10"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M36 12V6s-1.5 3-6 0s-6 0-6 0v6s1.5-3 6 0s6 0 6 0Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 44V31h-8v13"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M18 44h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSchool0)"/>`),
		g.Group(children),
	)
}