package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineUpRightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.832 2.012a1.95 1.95 0 0 1 2.154 2.154l-1.34 12.061c-.18 1.625-2.161 2.32-3.317 1.164l-1.267-1.267l-5.304 5.303a1.95 1.95 0 0 1-2.758 0l-5.429-5.429a1.95 1.95 0 0 1 0-2.758l5.304-5.303l-1.268-1.268c-1.156-1.156-.461-3.136 1.164-3.317l12.061-1.34Z"/>`),
		g.Group(children),
	)
}