package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nosurfreddit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 15.1a8.85 8.85 0 1 0 15 9.41a8.85 8.85 0 0 1 15 9.41A17.69 17.69 0 1 1 9 15.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.19 16.09a4.36 4.36 0 1 0-8.08 3.3a4.36 4.36 0 1 1-8.11 3.3a8.73 8.73 0 1 1 16.16-6.6"/>`),
		g.Group(children),
	)
}