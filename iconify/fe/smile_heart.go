package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feSmileHeart0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feSmileHeart1" fill="currentColor"><path id="feSmileHeart2" d="M10 22a8 8 0 1 1 0-16a8 8 0 0 1 0 16Zm0-2a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm3-5a3 3 0 0 1-6 0h6Zm-5-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm6.625-5c-.827-.18-3.375-1.59-3.375-4.125a1.875 1.875 0 0 1 3.375-1.125A1.875 1.875 0 0 1 22 3.875C22 6.41 19.452 7.82 18.625 8Z"/></g></g>`),
		g.Group(children),
	)
}