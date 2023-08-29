package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v20h9.586L16 29.414L19.414 26H29V6H3zm2 2h22v16h-8.414L16 26.586L13.414 24H5V8zm10 3v4h-4v2h4v4h2v-4h4v-2h-4v-4h-2z"/>`),
		g.Group(children),
	)
}