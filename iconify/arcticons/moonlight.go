package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5V21m0 6v18.5M45.5 24H27m-6 0H2.5M39.2 8.8L26.12 21.88m-4.24 4.24L8.8 39.2m30.4 0L26.12 26.12m-4.24-4.24L8.8 8.8"/>`),
		g.Group(children),
	)
}