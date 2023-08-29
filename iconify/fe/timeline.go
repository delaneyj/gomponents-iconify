package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feTimeline0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feTimeline1" fill="currentColor"><path id="feTimeline2" d="M9.17 17a3.001 3.001 0 0 1 5.66 0H20l1 1l-1 1h-5.17a3.001 3.001 0 0 1-5.66 0H3l1-1l-1-1h6.17ZM12 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm2-7l-2 2l-2-2H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-3ZM7 5v5h4l1 1l1-1h4V5H7Z"/></g></g>`),
		g.Group(children),
	)
}