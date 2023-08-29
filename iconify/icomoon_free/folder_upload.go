package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4H9zM8 7.5l3.5 3.5H9v4H7v-4H4.5L8 7.5z"/>`),
		g.Group(children),
	)
}