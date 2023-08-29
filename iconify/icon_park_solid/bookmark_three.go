package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBookmarkThree0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M7 9a3 3 0 0 1 3-3h31v36H10a3 3 0 0 1-3-3V9Z"/><path stroke="#000" stroke-linecap="round" d="M7 34h34"/><path stroke="#fff" stroke-linecap="round" d="M7 30v8m34-8v8"/><path fill="#000" stroke="#000" d="M15 6h10v20l-5-4l-5 4V6Z"/><path stroke="#fff" stroke-linecap="round" d="M11 6h18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBookmarkThree0)"/>`),
		g.Group(children),
	)
}