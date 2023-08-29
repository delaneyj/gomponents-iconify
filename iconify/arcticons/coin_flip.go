package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoinFlip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.938 35.5a8.326 8.326 0 0 1-.883-5.223s-3.53-1.36-4.119-2.096s.441-1.802.368-4.046"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.444 22.995c.478.98-2.207 1.324-2.39.147s2.06-2.427 2.17-3.751s1.688-7.503 8.275-6.842c7.687.773 6.804 8.607 6.694 10.041a2.911 2.911 0 0 1-1.545 4.819a8.179 8.179 0 0 0 .037 4.634"/>`),
		g.Group(children),
	)
}