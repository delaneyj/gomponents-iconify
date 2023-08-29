package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Justice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M262 247q16 16 30 0l90-89q13-15 0-30q-15-15-30 0l-90-90q16-14 0-29q-14-16-29 0l-90 87q-15 15 0 30q14 14 30 0l30 30l-60 59q-15-13-30 0L9 320q-15 15 0 30l27 34q16 16 30 0l105-105q6-6 6-14.5t-6-14.5l59-60l30 30q-12 12 2 27zm-29-119l-30-30l30-30l89 90l-30 30l-30-30z"/>`),
		g.Group(children),
	)
}