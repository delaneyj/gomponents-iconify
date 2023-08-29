package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallAmericanFootball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 9l-6 6m1-3l2 2m0-4l2 2m-6 9a5 5 0 0 0-5-5"/><path d="M16 3C8.82 3 3 8.82 3 16a5 5 0 0 0 5 5c7.18 0 13-5.82 13-13a5 5 0 0 0-5-5"/><path d="M16 3a5 5 0 0 0 5 5"/></g>`),
		g.Group(children),
	)
}