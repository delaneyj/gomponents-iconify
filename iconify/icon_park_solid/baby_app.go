package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBabyApp0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M43 26v6a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V12a2 2 0 0 1 2-2h9"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 40h40"/><circle cx="30" cy="17" r="9" fill="#fff" stroke="#fff" stroke-width="4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 18v-2m22 2v-2M30 8c-.167-1-1.2-3.2-4-4m4 4c.083-1 .6-3.2 2-4"/><circle cx="33" cy="16" r="2" fill="#000"/><circle cx="27" cy="16" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBabyApp0)"/>`),
		g.Group(children),
	)
}