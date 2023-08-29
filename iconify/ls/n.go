package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func N(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M72 128v64c42-40 109-64 171-64s109 24 151 64c44 42 72 102 72 168v294h-72V356c-2-87-64-157-151-157c-88 0-169 70-171 157v298H0V128h72z"/>`),
		g.Group(children),
	)
}