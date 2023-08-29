package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HolySword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHolySword0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m17 13l7-9l7 9l-5 26h-4l-5-26Z"/><path d="M17 39h14m-7 0v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHolySword0)"/>`),
		g.Group(children),
	)
}