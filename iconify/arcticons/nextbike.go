package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nextbike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="36.25" cy="27.48" r="7.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.25 23.61a7.25 7.25 0 1 0 13.38 3.87c4.13-.91 4.58-6.54 7-8c1.85-1.15 3.44-2.11 4.48-2.73a1.3 1.3 0 0 0 .24-2a7.18 7.18 0 0 0-2-1.43"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.66 24c-1.66-1.14-2.14-6.87-5.15-8.68S8.22 13.08 4.5 14a32.7 32.7 0 0 1 2.64 2.19l-1 1.05a5 5 0 0 0 2.48 1.53c-.43 2 2.57 6.62 8.15 6.62c3.84-.01 4.89-1.39 4.89-1.39Z"/>`),
		g.Group(children),
	)
}