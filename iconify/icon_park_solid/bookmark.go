package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBookmark0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M10 44a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v36a2 2 0 0 1-2 2H10Z"/><path fill="#000" fill-rule="evenodd" stroke="#000" stroke-linecap="round" d="M21 22V4h12v18l-6-6.273L21 22Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" d="M10 4h28"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBookmark0)"/>`),
		g.Group(children),
	)
}