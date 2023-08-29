package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallVolleyball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M12 12a8 8 0 0 0 8 4M7.5 13.5A12 12 0 0 0 16 20"/><path d="M12 12a8 8 0 0 0-7.464 4.928m8.415-9.575a12 12 0 0 0-9.88 4.111"/><path d="M12 12a8 8 0 0 0-.536-8.928m4.085 12.075a12 12 0 0 0 1.38-10.611"/></g>`),
		g.Group(children),
	)
}