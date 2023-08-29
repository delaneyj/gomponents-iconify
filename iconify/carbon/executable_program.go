package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExecutableProgram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m25 21l7 5l-7 5V21z"/><path fill="currentColor" d="m20.17 19l-2.59 2.59L19 23l4-4l-4-4l-1.42 1.41L20.17 19zm-8.34 0l2.59-2.59L13 15l-4 4l4 4l1.42-1.41L11.83 19z"/><circle cx="9" cy="8" r="1" fill="currentColor"/><circle cx="6" cy="8" r="1" fill="currentColor"/><path fill="currentColor" d="M21 26H4V12h24v7h2V6c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v20c0 1.103.897 2 2 2h17v-2ZM4 6h24v4H4V6Z"/>`),
		g.Group(children),
	)
}