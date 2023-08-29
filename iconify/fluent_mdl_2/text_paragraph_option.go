package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextParagraphOption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m308 640l-48 128H157l240-640h102l240 640H636l-48-128H308zm140-375l-92 247h184l-92-247zm64 1405l163-163l90 90l-317 317l-317-317l90-90l163 163V896h128v774zm512-1542h896v128h-896V128zm0 640V640h896v128h-896zm0 512v-128h896v128h-896zm0 512v-128h896v128h-896z"/>`),
		g.Group(children),
	)
}