package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v20h9.586L16 29.414L19.414 26H29V6zm2 2h22v16h-8.414L16 26.586L13.414 24H5zm4 3v2h14v-2zm0 4v2h14v-2zm0 4v2h10v-2z"/>`),
		g.Group(children),
	)
}