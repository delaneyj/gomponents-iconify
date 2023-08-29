package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Andotp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.48a16.12 16.12 0 0 1 15.83 13.13H8.17A16.11 16.11 0 0 1 24 8.48Zm0 18.18a2.59 2.59 0 0 1 1.43 4.75l1.43 3.88h-5.72l1.43-3.88A2.59 2.59 0 0 1 24 26.66Z"/>`),
		g.Group(children),
	)
}