package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M36 56v144a4 4 0 0 1-8 0V56a4 4 0 0 1 8 0Zm200 40v64a12 12 0 0 1-12 12H72a12 12 0 0 1-12-12V96a12 12 0 0 1 12-12h152a12 12 0 0 1 12 12Zm-8 0a4 4 0 0 0-4-4H72a4 4 0 0 0-4 4v64a4 4 0 0 0 4 4h152a4 4 0 0 0 4-4Z"/>`),
		g.Group(children),
	)
}