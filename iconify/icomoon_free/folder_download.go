package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 4L7 2H0v13h16V4H9zm-1 9.5L4.5 10H7V6h2v4h2.5L8 13.5z"/>`),
		g.Group(children),
	)
}