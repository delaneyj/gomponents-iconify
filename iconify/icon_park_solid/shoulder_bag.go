package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoulderBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShoulderBag0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M4.682 23.564A3 3 0 0 1 7.63 20h32.74a3 3 0 0 1 2.947 3.564l-3.062 16A3 3 0 0 1 37.309 42H10.691a3 3 0 0 1-2.947-2.436l-3.062-16Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m24 6l-9 14h18L24 6Z"/><circle cx="24" cy="31" r="3" fill="#000" stroke="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShoulderBag0)"/>`),
		g.Group(children),
	)
}