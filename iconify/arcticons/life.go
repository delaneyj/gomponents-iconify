package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Life(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.069 3.5c13.06 0 20.752 9.077 20.752 22.137c0 23.26-33.688 26.27-33.688 1.931c0-14.289 8.761-16.06 12.81-16.06c7.951 0 12.492 5.8 12.492 14.507c0 11.855-16.937 16.775-16.937.958c0-5.06.847-8.085 4.948-8.085c2.969 0 3.168 4.03 3.224 6.246C28.274 49 3.18 47.436 3.18 24.92c0-12.671 9.591-21.42 20.89-21.42Z"/>`),
		g.Group(children),
	)
}