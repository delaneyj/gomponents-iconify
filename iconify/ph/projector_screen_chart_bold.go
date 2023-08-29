package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenChartBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M104 128v8a12 12 0 0 1-24 0v-8a12 12 0 0 1 24 0Zm24-16a12 12 0 0 0-12 12v12a12 12 0 0 0 24 0v-12a12 12 0 0 0-12-12Zm36-4a12 12 0 0 0-12 12v16a12 12 0 0 0 24 0v-16a12 12 0 0 0-12-12Zm56-16.4V164h4a12 12 0 0 1 0 24h-84v23.22a24 24 0 1 1-24 0V188H32a12 12 0 0 1 0-24h4V91.6A20 20 0 0 1 20 72V48a20 20 0 0 1 20-20h176a20 20 0 0 1 20 20v24a20 20 0 0 1-16 19.6ZM44 68h168V52H44Zm152 96V92H60v72Z"/>`),
		g.Group(children),
	)
}