package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14v1.17L17.17 12L14 8.83V10H6v4z" opacity=".3"/><path fill="currentColor" d="m20 12l-8-8v4H4v8h8v4l8-8zM6 14v-4h8V8.83L17.17 12L14 15.17V14H6z"/>`),
		g.Group(children),
	)
}