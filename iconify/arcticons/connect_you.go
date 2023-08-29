package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.97 39.04v-3.719c2.974-2.974 5.467-3.918 9.265-3.918h0c3.798 0 6.29.944 9.265 3.918v3.719H24.97Z"/><circle cx="34.235" cy="22.893" r="4.975" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.97 39.04H4.5v-6.289c5.03-5.03 9.245-6.625 15.668-6.625h0c5.765 0 9.752 1.286 14.146 5.188"/><circle cx="20.168" cy="15.809" r="6.848" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}