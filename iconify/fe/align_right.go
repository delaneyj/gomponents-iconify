package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAlignRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignRight1" fill="currentColor"><path id="feAlignRight2" d="M17 21a1 1 0 0 1-2 0V3a1 1 0 0 1 2 0v18ZM11 5a2 2 0 1 1 0 4H5a2 2 0 1 1 0-4h6Zm0 7a2 2 0 1 1 0 4H7a2 2 0 1 1 0-4h4Z"/></g></g>`),
		g.Group(children),
	)
}