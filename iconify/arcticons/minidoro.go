package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minidoro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.799 17.881c5.61-6.238 17.88-7.339 19.605 5.71c1.112 8.416-7.527 16.623-19.605 16.78c-9.573.124-19.729-7.755-19.285-17.013c.63-13.136 16.374-11.578 19.285-5.477Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.799 17.881c1.754-3.465 4.324-6.268 9.322-7.107"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.944 7.628c6.974 2.334 6.823 5.854 5.855 10.254c-1.112-3.314-3.575-5.951-9.176-7.02m9.176 9.614v6.902m0 0l6.321 7.02"/>`),
		g.Group(children),
	)
}