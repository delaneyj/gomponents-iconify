package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M38 56v144a6 6 0 0 1-12 0V56a6 6 0 0 1 12 0Zm200 40v64a14 14 0 0 1-14 14H72a14 14 0 0 1-14-14V96a14 14 0 0 1 14-14h152a14 14 0 0 1 14 14Zm-12 0a2 2 0 0 0-2-2H72a2 2 0 0 0-2 2v64a2 2 0 0 0 2 2h152a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}