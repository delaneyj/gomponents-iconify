package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Norisbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.15 2.58C17.93 4.448 8.257 13.448 6.407 20.387M26.514 2.646c-4.774 2.45-13.505 12.4-13.827 21.99M30.168 3.398c-5.327 3.03-12.406 13.247-11.2 25.485"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.063 4.499c-5.88 3.609-10.549 13.744-7.815 28.632"/>`),
		g.Group(children),
	)
}