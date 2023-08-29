package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2v17l-4-4l-4 4V2c0-.553.585-1.02 1-1h6c.689-.02 1 .447 1 1z"/>`),
		g.Group(children),
	)
}