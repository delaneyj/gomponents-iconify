package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Totem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18h4V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v15zm6 .268A2 2 0 0 1 15 20H5a2 2 0 0 1 1-1.732V3a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v15.268zM9 9h2v2H9V9zm1-4a1 1 0 1 1 0-2a1 1 0 0 1 0 2zM1.286 6h17.433a1 1 0 0 1 .97 1.243l-1.31 5.242A2 2 0 0 1 16.439 14H3.588a2 2 0 0 1-1.939-1.507L.317 7.246A1 1 0 0 1 1.286 6zm1.286 2l1.017 4h12.85l1-4H2.571z"/>`),
		g.Group(children),
	)
}