package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIron0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M4 40h40l-2-16H20c-8.837 0-16 7.163-16 16Z"/><path stroke="#fff" d="M16 8h24l2 16"/><path stroke="#000" d="M17 32h2m6 0h2m6 0h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIron0)"/>`),
		g.Group(children),
	)
}