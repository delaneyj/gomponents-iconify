package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Booking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.55 5.09h-23.1A7 7 0 0 0 5.5 12v30.09h30.05a7 7 0 0 0 7-6.94V12a7 7 0 0 0-7-6.91ZM32.88 34a2.48 2.48 0 1 1 2.47-2.48A2.47 2.47 0 0 1 32.88 34Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M22.89 24a5 5 0 0 1 0 10h-8.25V14h8.25a5 5 0 0 1 0 10Zm0 0h-8.25"/>`),
		g.Group(children),
	)
}