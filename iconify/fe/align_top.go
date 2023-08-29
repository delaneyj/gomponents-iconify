package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAlignTop0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignTop1" fill="currentColor"><path id="feAlignTop2" d="M4.8 3h14.4c.442 0 .8.448.8 1s-.358 1-.8 1H4.8C4.358 5 4 4.552 4 4s.358-1 .8-1ZM7 9a2 2 0 1 1 4 0v10a2 2 0 1 1-4 0V9Zm6 0a2 2 0 1 1 4 0v6a2 2 0 1 1-4 0V9Z"/></g></g>`),
		g.Group(children),
	)
}