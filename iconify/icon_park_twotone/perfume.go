package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Perfume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPerfume0"><g fill="none"><rect width="16" height="10" x="16" y="4" fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="3"/><rect width="36" height="24" x="6" y="20" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="12"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 33c3.042-3.867 12-3 18-1s14 5 18 0"/><path fill="#fff" d="M25 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPerfume0)"/>`),
		g.Group(children),
	)
}