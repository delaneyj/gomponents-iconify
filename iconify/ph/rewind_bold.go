package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M225.7 54.46a20 20 0 0 0-20.33.66L132 101.85v-30a19.91 19.91 0 0 0-30.63-16.72l-88.18 56.16a19.79 19.79 0 0 0 0 33.42l88.18 56.17A19.91 19.91 0 0 0 132 184.16v-30l73.37 46.73A19.91 19.91 0 0 0 236 184.16V71.84a19.84 19.84 0 0 0-10.3-17.38ZM108 176.64L31.63 128L108 79.36Zm104 0L135.63 128L212 79.36Z"/>`),
		g.Group(children),
	)
}