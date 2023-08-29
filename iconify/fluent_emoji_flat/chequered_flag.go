package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChequeredFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#F4F4F4" d="m3 7l-1 3v4l1 2l-1 2v4l1 3l3 1h4l2-1l2 1h4l2-1l2 1h4l3-1l1-3v-4l-1-2l1-2v-4l-1-3l-3-1h-4l-2 1l-2-1h-4l-2 1l-2-1H6L3 7Z"/><path fill="#000" d="M18 14v4h4v-4h-4Z"/><path fill="#000" d="M6 6H3a1 1 0 0 0-1 1v3h4v4H2v4h4v4H2v3a1 1 0 0 0 1 1h3v-4h4v4h4v-4h4v4h4v-4h4v4h3a1 1 0 0 0 1-1v-3h-4v-4h4v-4h-4v-4h4V7a1 1 0 0 0-1-1h-3v4h-4V6h-4v4h-4V6h-4v4H6V6Zm4 8v-4h4v4h-4Zm0 4h4v-4h4v-4h4v4h4v4h-4v4h-4v-4h-4v4h-4v-4Zm0 0H6v-4h4v4Z"/></g>`),
		g.Group(children),
	)
}