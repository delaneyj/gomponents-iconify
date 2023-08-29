package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Privatelock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.2 4.86L6.69 11.25V27C6.69 35.44 24 43.5 24 43.5S41.31 35.44 41.31 27V11.25L25.8 4.86a4.68 4.68 0 0 0-3.6 0Z"/><rect width="17.39" height="13.4" x="15.31" y="19.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.05"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.25 19.26v-2.41a5.75 5.75 0 0 1 11.5 0v2.41"/><circle cx="24" cy="25.96" r="2.23" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}