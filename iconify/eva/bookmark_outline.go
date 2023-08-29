package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBookmarkOutline0"><g id="evaBookmarkOutline1"><path id="evaBookmarkOutline2" fill="currentColor" d="M6.09 21.06a1 1 0 0 1-1-1L4.94 5.4a2.26 2.26 0 0 1 2.18-2.35L16.71 3a2.27 2.27 0 0 1 2.23 2.31l.14 14.66a1 1 0 0 1-.49.87a1 1 0 0 1-1 0l-5.7-3.16l-5.29 3.23a1.2 1.2 0 0 1-.51.15Zm5.76-5.55a1.11 1.11 0 0 1 .5.12l4.71 2.61l-.12-12.95c0-.2-.13-.34-.21-.33l-9.6.09c-.08 0-.19.13-.19.33l.12 12.9l4.28-2.63a1.06 1.06 0 0 1 .51-.14Z"/></g></g>`),
		g.Group(children),
	)
}