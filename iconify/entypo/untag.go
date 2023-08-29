package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Untag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m1 14.742l4.945-.709L5.239 19l5.962-5.985l-4.069-4.429L1 14.742zm17.664-9.221c.391-.393.5-.945 0-1.419l-2.826-2.839c-.279-.308-1.021-.392-1.412 0l-3.766 3.78l4.068 4.429l3.936-3.951zm.042 9.772l-14.001-14a.999.999 0 1 0-1.414 1.414l14.001 14a.996.996 0 0 0 1.414 0a.999.999 0 0 0 0-1.414z"/>`),
		g.Group(children),
	)
}