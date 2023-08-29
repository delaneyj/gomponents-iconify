package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxlSteam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M8.21 17.32L7 16.8a2.13 2.13 0 1 0 1.17-2.93l1.28.53a1.58 1.58 0 0 1-1.22 2.92z" fill="currentColor"/><path d="M12 2a10 10 0 0 0-10 9.34l5.38 2.21a2.31 2.31 0 0 1 .47-.24A2.62 2.62 0 0 1 9 13.1l2.44-3.56a3.8 3.8 0 1 1 3.8 3.8h-.08l-3.51 2.5a2.77 2.77 0 0 1-5.47.68l-3.77-1.6A10 10 0 1 0 12 2z" fill="currentColor"/><path d="M17.79 9.5a2.53 2.53 0 1 0-2.53 2.5a2.54 2.54 0 0 0 2.53-2.5zm-4.42 0a1.9 1.9 0 1 1 1.9 1.91a1.9 1.9 0 0 1-1.9-1.92z" fill="currentColor"/>`),
		g.Group(children),
	)
}