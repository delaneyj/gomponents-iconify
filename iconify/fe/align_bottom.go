package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAlignBottom0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAlignBottom1" fill="currentColor"><path id="feAlignBottom2" d="M4.8 19h14.4c.442 0 .8.448.8 1s-.358 1-.8 1H4.8c-.442 0-.8-.448-.8-1s.358-1 .8-1Zm6.2-4a2 2 0 1 1-4 0V5a2 2 0 1 1 4 0v10Zm6 0a2 2 0 1 1-4 0V9a2 2 0 1 1 4 0v6Z"/></g></g>`),
		g.Group(children),
	)
}