package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GirlOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGirlOne0"><g fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="11" r="7"/><path d="M27 24h-6L8 44h32L27 24Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGirlOne0)"/>`),
		g.Group(children),
	)
}