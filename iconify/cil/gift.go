package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 128H371.833A71.962 71.962 0 0 0 256 42.815A71.962 71.962 0 0 0 140.167 128H40a24.028 24.028 0 0 0-24 24v321.556C16 485.932 26.767 496 40 496h432c13.233 0 24-10.068 24-22.444V152a24.028 24.028 0 0 0-24-24ZM312 48a40 40 0 0 1 0 80h-40V88a40.045 40.045 0 0 1 40-40ZM160 88a40 40 0 0 1 80 0v40h-40a40.045 40.045 0 0 1-40-40ZM48 464V256h144v208Zm176 0V256h64v208Zm240 0H320V256h144ZM48 224v-64h416v64Z"/>`),
		g.Group(children),
	)
}