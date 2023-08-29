package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFile0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFile1" fill="currentColor"><path id="feFile2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4Z"/></g></g>`),
		g.Group(children),
	)
}