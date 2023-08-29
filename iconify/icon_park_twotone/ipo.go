package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIpo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 18.313V7h40v11.313a5.5 5.5 0 0 0-3.636 5.187A5.5 5.5 0 0 0 44 28.687V40H4V28.687A5.5 5.5 0 0 0 7.636 23.5A5.5 5.5 0 0 0 4 18.313Z"/><path stroke-linecap="round" d="M13 18v11m5-11v11"/><path fill="#555" d="M18 18h3a3 3 0 1 1 0 6h-3v-6Z"/><ellipse cx="32" cy="24" fill="#555" rx="3" ry="5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIpo0)"/>`),
		g.Group(children),
	)
}