package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4H9zm2 7H9v2H7v-2H5V9h2V7h2v2h2v2z"/>`),
		g.Group(children),
	)
}