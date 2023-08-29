package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M112 176h32v-32h80v224h-32v32h128v-32h-32V144h80v32h32v-64H112v64z"/>`),
		g.Group(children),
	)
}