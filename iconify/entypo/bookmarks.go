package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 0h-4a1 1 0 0 0-1 1l.023.222c1.102 0 2 .897 2 2v11.359L13 13.4l3 3.6V1a1 1 0 0 0-1-1zM9.023 3H5a1 1 0 0 0-1 1v16l3-3.6l3 3.6V4c0-.553-.424-1-.977-1z"/>`),
		g.Group(children),
	)
}