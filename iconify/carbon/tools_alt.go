package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23.1 16l6.3-6.3c.8-.8.8-2 0-2.8l-4.2-4.2c-.8-.8-2-.8-2.8 0L16 8.9L9.7 2.6c-.8-.8-2-.8-2.8 0L2.6 6.8c-.8.8-.8 2 0 2.8L8.9 16L2 22.9V30h7.1l6.9-6.9l6.3 6.3c.8.8 2 .8 2.8 0l4.2-4.2c.8-.8.8-2 0-2.8L23.1 16zm.7-12L28 8.2l-6.3 6.3l-4.2-4.2L23.8 4zM8.2 28H4v-4.2l6.3-6.3l4.2 4.2L8.2 28zm15.6 0L4 8.2L8.2 4l3.5 3.5l-2.1 2.1L11 11l2.1-2.1l4.2 4.2l-2.1 2.1l1.4 1.4l2.1-2.1l4.2 4.2l-1.9 2.2l1.4 1.4l2.1-2.1l3.5 3.5l-4.2 4.3z"/>`),
		g.Group(children),
	)
}