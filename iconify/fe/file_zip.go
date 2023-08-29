package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFileZip0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFileZip1" fill="currentColor"><path id="feFileZip2" d="M6 2h10l4 4v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2Zm9.172 2H6v16h12V6.828h-2.828V4ZM8 4h2v2H8V4Zm2 2h2v2h-2V6ZM8 8h2v2H8V8Zm2 2h2v2h-2v-2Zm-1 4l1-1l1 1v4H9v-4Z"/></g></g>`),
		g.Group(children),
	)
}