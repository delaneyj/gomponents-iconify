package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Application(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 20H3v-1H2V3h1V2h16v1h1v16h-1ZM18 6V4H4v2Zm0 12V8H4v10Z"/>`),
		g.Group(children),
	)
}