package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkFilledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilBookmarkFilledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" d="m13 16.467l-3.755 4.2C8.634 21.35 7.5 20.918 7.5 20V6a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v14c0 .918-1.133 1.35-1.745.666L13 16.468Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilBookmarkFilledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}