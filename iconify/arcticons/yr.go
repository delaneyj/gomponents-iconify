package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.5 31.7V16.3h5.1c2.9 0 5.2 2.3 5.2 5.2s-2.3 5.2-5.2 5.2h-5.1m5 0l5.1 5M21 23.3v4.5c0 2.1-1.7 3.9-3.9 3.9h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 16.3v7.1c0 2.1-1.7 3.9-3.9 3.9h0c-2.1 0-3.9-1.7-3.9-3.9h0v-7.1"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}