package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="m24 10l-12 8l12 8l12-8l-12-8Z"/><path d="m24 10l12 8l5-7l-11-7l-6 6Zm0 0l-12 8l-5-7l11-7l6 6Zm19 12l-7-4l-12 8l7 5l12-9ZM5 22l7-4l12 8l-7 5l-12-9Z"/><path stroke-linecap="round" d="M36 28v9l-12 7l-12-7v-9"/></g>`),
		g.Group(children),
	)
}