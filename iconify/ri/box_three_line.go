package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM5.494 7.078L12 10.844l6.506-3.766L12 3.31L5.494 7.078ZM4.5 8.813v7.534L11 20.11v-7.534L4.5 8.813ZM13 20.11l6.5-3.763V8.813L13 12.576v7.534Z"/>`),
		g.Group(children),
	)
}