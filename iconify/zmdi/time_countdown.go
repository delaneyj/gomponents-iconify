package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimeCountdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 298.5q0-8.5 6-15t15-6.5t15 6.5t6 15t-6 15t-15 6.5t-15-6.5t-6-15zM171 0h21q80 0 136 56t56 136t-56 136t-136 56t-136-56T0 192q0-46 20.5-86.5T77 39v-1l145 145l-30 30L76 98q-33 41-33 94q0 62 43.5 105.5T192 341t105.5-43.5T341 192q0-56-36.5-98T213 44v41h-42V0zm149 192q0 9-6.5 15t-15 6t-15-6t-6.5-15t6.5-15t15-6t15 6t6.5 15zm-256 0q0-9 6.5-15t15-6t15 6t6.5 15t-6.5 15t-15 6t-15-6t-6.5-15z"/>`),
		g.Group(children),
	)
}