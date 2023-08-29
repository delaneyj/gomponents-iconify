package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Header(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M288 144h32v96H192v-96h32v-32H96v32h32v224H96v32h128v-32h-32v-96h128v96h-32v32h128v-32h-32V144h32v-32H288v32z"/>`),
		g.Group(children),
	)
}