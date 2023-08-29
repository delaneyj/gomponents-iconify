package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Runkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.97 14.797a3 3 0 0 1-1.47 3.02l-9 5.2a3 3 0 0 1-3 0l-9-5.2a3 3 0 0 1-1.47-3.02l1.32-7.2a3 3 0 0 1 .98-1.82a2.96 2.96 0 0 1 .49-.35l7.55-4.37a3.01 3.01 0 0 1 3.28.02l7.53 4.35c.1.05.19.1.28.17a3 3 0 0 1 1.19 2zm-9.54-10.77l-7.72 1.59c-.43.08-.51.64-.14.86l5.6 3.23c.23.13.5.07.63-.19l1.58-3.6l.53-1.22c.16-.35-.11-.75-.5-.67z"/>`),
		g.Group(children),
	)
}