package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M275.2 512L480 409.6l20.5-307.2l-225.3 61.4V512zM29.5 409.6L234.3 512V163.8L9 102.4l20.5 307.2zM254.8 0L9 61.4l245.8 61.4l245.8-61.4L254.8 0z"/>`),
		g.Group(children),
	)
}