package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 5v18h5v5.078L14.352 23H29V5zm2 2h22v14H13.648L10 23.918V21H5z"/>`),
		g.Group(children),
	)
}