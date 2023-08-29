package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMovie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFileMovie0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileMovie1" fill="currentColor"><path id="feFileMovie2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM12 13a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1v-3a1 1 0 0 1 1-1h3Zm1 2.5l3-1.5v3l-3-1.5Z"/></g></g>`),
		g.Group(children),
	)
}