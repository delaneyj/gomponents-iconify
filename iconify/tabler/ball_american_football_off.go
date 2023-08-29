package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallAmericanFootballOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 9l-1 1m-2 2l-3 3m1-3l2 2m-4 7a5 5 0 0 0-5-5"/><path d="M6.813 6.802A12.96 12.96 0 0 0 3 16a5 5 0 0 0 5 5a12.96 12.96 0 0 0 9.186-3.801m1.789-2.227A12.94 12.94 0 0 0 21 8a5 5 0 0 0-5-5a12.94 12.94 0 0 0-6.967 2.022"/><path d="M16 3a5 5 0 0 0 5 5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}