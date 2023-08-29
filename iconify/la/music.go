package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27 3.781l-1.188.25l-16 3L9 7.156v13.407A3.93 3.93 0 0 0 7 20c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4V12.812l14-2.624v7.374A3.93 3.93 0 0 0 23 17c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4zm-2 2.406v2l-14 2.626v-2zM23 19c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2zM7 22c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2z"/>`),
		g.Group(children),
	)
}