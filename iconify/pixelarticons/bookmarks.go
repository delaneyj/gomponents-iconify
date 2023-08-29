package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18V2H7v2h12v14h2zM5 6H3v16h4v-2h2v-2h2v2h2v2h4V6H5zm8 14v-2h-2v-2H9v2H7v2H5V8h10v12h-2z"/>`),
		g.Group(children),
	)
}