package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 2v14l5-5l5 5V2zm8-2H2v14l1-1V1h9z"/>`),
		g.Group(children),
	)
}