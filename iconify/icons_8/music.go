package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27 3.78l-1.188.25l-16 3L9 7.157v13.407A3.893 3.893 0 0 0 7 20c-2.197 0-4 1.803-4 4s1.803 4 4 4s4-1.803 4-4V12.812l14-2.624v7.374A3.93 3.93 0 0 0 23 17c-2.197 0-4 1.803-4 4s1.803 4 4 4s4-1.803 4-4V3.78zm-2 2.41v2l-14 2.624v-2L25 6.19zM23 19c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2zM7 22c1.116 0 2 .884 2 2s-.884 2-2 2s-2-.884-2-2s.884-2 2-2z"/>`),
		g.Group(children),
	)
}