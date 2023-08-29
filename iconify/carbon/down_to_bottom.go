package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownToBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 18L6 8l1.4-1.4l8.6 8.6l8.6-8.6L26 8zM4 22h24v2H4z"/>`),
		g.Group(children),
	)
}