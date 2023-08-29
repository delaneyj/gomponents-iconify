package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feExpand0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feExpand1" fill="currentColor"><path id="feExpand2" d="M14 4h6v6l-2-2l-4 4l-2-2l4-4l-2-2Zm-4 16H4v-6l2 2l4-4l2 2l-4 4l2 2Z"/></g></g>`),
		g.Group(children),
	)
}