package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 12.5a.47.47 0 0 1-.35-.15l-4.5-4.5C1.06 7.76 1 7.63 1 7.5s.05-.26.15-.35l4.5-4.5c.2-.2.51-.2.71 0c.2.2.2.51 0 .71L2.21 7.5l4.15 4.15c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M13.5 14c-.28 0-.5-.22-.5-.5v-3A2.5 2.5 0 0 0 10.5 8H2.7c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h7.8c1.93 0 3.5 1.57 3.5 3.5v3c0 .28-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}