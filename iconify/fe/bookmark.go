package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feBookmark0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feBookmark1" fill="currentColor" fill-rule="nonzero"><path id="feBookmark2" d="M18 4H6v14.764l6-3l6 3V4ZM6 2h12a2 2 0 0 1 2 2v18l-8-4l-8 4V4a2 2 0 0 1 2-2Zm8 4h2v6h-2V6Z"/></g></g>`),
		g.Group(children),
	)
}