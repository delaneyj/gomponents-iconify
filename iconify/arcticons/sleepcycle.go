package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleepcycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05" d="M24 42.78a17.21 17.21 0 1 1 17.2-17.2A17.2 17.2 0 0 1 24 42.78Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.05" d="M27.78 8.8c3.85-3.93 9.47-3.67 13.38.68c2.5 2.76 3.76 8.56-.42 12.27m-33.55 0C3 18 4.27 12.24 6.76 9.48c3.92-4.35 9.53-4.61 13.38-.68m2.16 20.31l10.26-10.26Zm0 0l-6.85-6.85"/>`),
		g.Group(children),
	)
}