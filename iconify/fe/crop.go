package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCrop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCrop1" fill="currentColor"><path id="feCrop2" d="M16 20v-2H8c-1.1 0-2-.9-2-2V8H4a1 1 0 1 1 0-2h2V4a1 1 0 1 1 2 0v2h9l2-2l1 1l-2 2v9h2a1 1 0 0 1 0 2h-2v2a1 1 0 0 1-2 0Zm0-4V9l-7 7h7ZM8 8v7l7-7H8Z"/></g></g>`),
		g.Group(children),
	)
}