package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTextAlignRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTextAlignRight1" fill="currentColor"><path id="feTextAlignRight2" d="M19 7H5a1 1 0 1 1 0-2h14a1 1 0 0 1 0 2Zm0 4H9a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Zm0 4H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2Zm0 4H9a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2Z"/></g></g>`),
		g.Group(children),
	)
}