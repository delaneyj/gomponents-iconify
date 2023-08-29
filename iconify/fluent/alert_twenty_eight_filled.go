package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.466 22.002a3.502 3.502 0 0 1-6.932 0h6.932ZM14 3a8.501 8.501 0 0 1 8.497 8.246v.255h.004v4.112l1.414 3.644c.038.099.064.201.077.306l.01.157a1.28 1.28 0 0 1-1.15 1.274l-.13.006H5.275a1.28 1.28 0 0 1-1.235-1.62l.042-.124l1.416-3.644v-4.11A8.501 8.501 0 0 1 14 3Z"/>`),
		g.Group(children),
	)
}