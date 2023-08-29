package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAlignLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignLeft1" fill="currentColor"><path id="feAlignLeft2" d="M7 21a1 1 0 0 1-2 0V3a1 1 0 1 1 2 0v18Zm4-12a2 2 0 1 1 0-4h6a2 2 0 1 1 0 4h-6Zm0 3h4a2 2 0 1 1 0 4h-4a2 2 0 1 1 0-4Z"/></g></g>`),
		g.Group(children),
	)
}