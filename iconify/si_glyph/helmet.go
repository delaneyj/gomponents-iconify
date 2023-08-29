package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Helmet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.706 8.697h-.124c-.344-3.754-3.457-6.693-7.253-6.693c-4.026 0-7.288 3.305-7.288 7.383c0 2.579 1.192 4.589 5.218 4.589c3.146 0 1.98-2.803 5.315-4.026h4.132c.702 0 1.272-.271 1.272-.601v-.051c.001-.331-.57-.601-1.272-.601zm-9.644 2.345H4.937V9.958h1.125v1.084z"/>`),
		g.Group(children),
	)
}