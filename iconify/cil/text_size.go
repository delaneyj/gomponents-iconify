package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 184h32v-48h96v232h-40v32h144v-32h-40V136h96v48h32v-80H176v80z"/><path fill="currentColor" d="M16 280h32v-32h56v152H72v32h112v-32h-32V248h64v32h32v-64H16v64z"/>`),
		g.Group(children),
	)
}