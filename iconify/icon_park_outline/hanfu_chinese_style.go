package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HanfuChineseStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m9 9l9-5h12l9 5l4 16l-8 4v15H13V29l-8-4L9 9Zm9-5l6 10.5"/><path d="m30 4l-6 10.5L13 29"/></g>`),
		g.Group(children),
	)
}