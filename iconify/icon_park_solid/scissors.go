package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSScissors0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" d="M19 19c2 3 2 7 0 10m17-15L21 24l15 10"/><circle cx="16" cy="16" r="4" stroke="#000"/><circle cx="16" cy="32" r="4" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSScissors0)"/>`),
		g.Group(children),
	)
}