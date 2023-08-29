package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSmilePlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSmilePlus1" fill="currentColor"><path id="feSmilePlus2" d="M10 22a8 8 0 1 1 0-16a8 8 0 0 1 0 16Zm0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3-5a3 3 0 0 1-6 0h6Zm-5-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm8-9h1a1 1 0 0 1 0 2h-1v1a1 1 0 0 1-2 0V6h-1a1 1 0 0 1 0-2h1V3a1 1 0 0 1 2 0v1Z"/></g></g>`),
		g.Group(children),
	)
}