package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func M(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M71 128v64c42-40 109-64 171-64s110 24 152 64c13 13 26 28 36 44c41-66 124-108 207-108c62 0 108 24 150 64c45 42 73 102 73 168v294h-73V356c-2-87-62-157-150-157s-170 71-171 159v296h-72V356c-3-87-64-157-152-157c-87 0-169 70-171 157v298H0V128h71z"/>`),
		g.Group(children),
	)
}