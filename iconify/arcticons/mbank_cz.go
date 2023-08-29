package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MbankCz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.22 22.13a6.134 6.134 0 0 1 5.993-6.232a6.134 6.134 0 0 1 5.993 6.232v9.971M13.22 15.898v16.204"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.206 22.13a5.998 5.998 0 1 1 11.986 0v9.971"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="6.25" d="M13.22 15.898h-2.403"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="6.25" d="M13.22 5.398v6.479m0 24.723v6.002M25.206 2.745v9.131m.02 24.724l-.02 8.793M37.154 6.994v4.883m0 24.723v4.406"/>`),
		g.Group(children),
	)
}