package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 20h-1.828l-4.465-4.465a1 1 0 0 0-1.414 0L6.828 20H5V7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v13ZM17 7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v10l2.879-2.879a3 3 0 0 1 4.242 0L17 17V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}