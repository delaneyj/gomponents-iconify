package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22h20V2H2v20Zm2-2v-5h4v5H4Zm6 0v-5h4v5h-4Zm6 0v-5h4v5h-4Zm4-7h-4V8h4v5Zm0-7H4V4h16v2ZM4 8h4v5H4V8Zm6 5V8h4v5h-4Z"/>`),
		g.Group(children),
	)
}