package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 6a4 4 0 0 0-4.608-3.953l-7 1.077A4 4 0 0 0 7 7.078v8.763a3.5 3.5 0 1 0 2 3.163V7.078A2 2 0 0 1 10.696 5.1l7-1.077A2 2 0 0 1 20 6.001v6.84a3.5 3.5 0 1 0 2 3.163V6.001Zm-2 10.004a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm-13 3a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}