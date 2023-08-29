package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopBookmarkCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M9.245 20.666L13 16.468l3.755 4.2c.612.684 1.745.251 1.745-.667V6a1 1 0 0 0-1-1h-9a1 1 0 0 0-1 1v14c0 .918 1.134 1.35 1.745.666Zm.255-3.285V7h7v10.381l-2.755-3.08a1 1 0 0 0-1.49 0L9.5 17.38Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopBookmarkCircleFilled0)"/></g>`),
		g.Group(children),
	)
}