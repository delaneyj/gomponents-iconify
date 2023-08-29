package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func H(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M72 0v292c42-40 109-64 171-64s110 24 152 64c44 42 71 102 71 168v294h-71V456c-3-87-64-157-152-157c-87 0-169 70-171 157v298H0V0h72z"/>`),
		g.Group(children),
	)
}