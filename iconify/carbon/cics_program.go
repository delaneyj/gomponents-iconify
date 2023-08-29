package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsProgram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m20.17 19l-2.59 2.59L19 23l4-4l-4-4l-1.42 1.41L20.17 19zm-8.34 0l2.59-2.59L13 15l-4 4l4 4l1.42-1.41L11.83 19z"/><circle cx="9" cy="8" r="1" fill="currentColor"/><circle cx="6" cy="8" r="1" fill="currentColor"/><path fill="currentColor" d="M28 4H4c-1.103 0-2 .897-2 2v20c0 1.103.897 2 2 2h24c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2Zm0 2v4H4V6h24ZM4 26V12h24v14H4Z"/>`),
		g.Group(children),
	)
}