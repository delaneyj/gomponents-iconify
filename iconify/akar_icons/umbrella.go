package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4c-4.915 0-8.91 3.378-8.999 8.817a.18.18 0 0 0 .182.183a.188.188 0 0 0 .17-.11C3.876 11.767 4.782 11.5 6 11.5c1.185 0 1.964.227 2.456 1.302c.054.12.172.198.304.198a.366.366 0 0 0 .326-.224C9.56 11.729 10.901 11.5 12 11.5M12 4c4.916 0 8.91 3.378 8.998 8.817a.18.18 0 0 1-.18.183a.188.188 0 0 1-.17-.11c-.525-1.123-1.43-1.39-2.648-1.39c-1.185 0-1.964.227-2.456 1.302a.336.336 0 0 1-.304.198a.366.366 0 0 1-.326-.224C14.44 11.729 13.099 11.5 12 11.5M12 4V2m0 9.5V20a2 2 0 1 1-4 0"/>`),
		g.Group(children),
	)
}