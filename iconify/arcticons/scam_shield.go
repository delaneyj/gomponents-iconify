package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScamShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c9.043-3.117 15.489-10.363 16.5-19.589a79.36 79.36 0 0 0-.071-12.027a2.541 2.541 0 0 0-2.468-2.366c-4.091-.126-8.846-.808-12.52-4.427a2.052 2.052 0 0 0-2.881 0c-3.675 3.619-8.43 4.301-12.52 4.427a2.541 2.541 0 0 0-2.468 2.366A79.36 79.36 0 0 0 7.5 23.911C8.511 33.137 14.957 40.383 24 43.5ZM24 33V15m-7.102 0h14.204m0 4.052V15m-14.204 4.052V15m2.68 18h8.844"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.9 23.349h3.203v3.203H27.9zm-11.002 0h3.203v3.203h-3.203z"/>`),
		g.Group(children),
	)
}