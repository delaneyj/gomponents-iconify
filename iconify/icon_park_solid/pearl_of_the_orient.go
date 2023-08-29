package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PearlOfTheOrient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPearlOfTheOrient0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="33" r="6" fill="#fff"/><circle cx="24" cy="13" r="3" fill="#fff"/><path stroke-linecap="round" d="M24 40v4m-2-28l-2 12m6-12l2 12m-9 9l-3 7m13-7l3 7M24 4v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPearlOfTheOrient0)"/>`),
		g.Group(children),
	)
}