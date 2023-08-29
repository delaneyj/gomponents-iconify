package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dsub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.31 33.36V13.89l17.25-2.39v17.9"/><circle cx="14.35" cy="33.36" r="3.96" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.61" cy="29.4" r="3.96" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24.8a10 10 0 0 0 9.25-6.2a10 10 0 0 0 18.5 0a10 10 0 0 0 9.25 6.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2H7.48Z"/>`),
		g.Group(children),
	)
}