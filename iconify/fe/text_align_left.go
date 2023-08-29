package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTextAlignLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextAlignLeft1" fill="currentColor"><path id="feTextAlignLeft2" d="M19 7H5a1 1 0 1 1 0-2h14a1 1 0 0 1 0 2Zm-4 4H5a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm4 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm-4 4H5a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/></g></g>`),
		g.Group(children),
	)
}