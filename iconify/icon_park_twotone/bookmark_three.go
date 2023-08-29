package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBookmarkThree0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M7 9a3 3 0 0 1 3-3h31v36H10a3 3 0 0 1-3-3V9Z"/><path stroke-linecap="round" d="M7 34h34M7 30v8m34-8v8"/><path fill="#555" d="M15 6h10v20l-5-4l-5 4V6Z"/><path stroke-linecap="round" d="M11 6h18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBookmarkThree0)"/>`),
		g.Group(children),
	)
}