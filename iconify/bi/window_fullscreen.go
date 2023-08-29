package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowFullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M3 3.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm1 .5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/><path d="M.5 1a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5v-13a.5.5 0 0 0-.5-.5H.5ZM1 5V2h14v3H1Zm0 1h14v8H1V6Z"/></g>`),
		g.Group(children),
	)
}