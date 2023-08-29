package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBookmarkOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M34 10V4H8v34l6-3"/><path fill="#fff" d="M14 44V10h26v34l-13-6.273L14 44Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBookmarkOne0)"/>`),
		g.Group(children),
	)
}