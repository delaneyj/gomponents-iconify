package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyoxide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.038" cy="9.886" r="1.363" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="25.188" cy="6.478" r=".977" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.288" cy="5.153" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.75 34.862l-8.498-6.356l8.498-6.355a4.85 4.85 0 0 0 .979-6.79a4.77 4.77 0 0 0-.36-.42a1.646 1.646 0 1 0-2.099-1.292a4.936 4.936 0 0 0-1.951-.185a2.594 2.594 0 1 0-2.729 1.18l-5.585 4.177v-.555a4.85 4.85 0 1 0-9.7 0v20.48a4.85 4.85 0 1 0 9.7 0v-.554l5.935 4.438a4.85 4.85 0 1 0 5.81-7.768Z"/>`),
		g.Group(children),
	)
}