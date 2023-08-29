package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 8.83L17.17 12L14 15.17V14H6v-4h8V8.83M12 4v4H4v8h8v4l8-8l-8-8z"/>`),
		g.Group(children),
	)
}