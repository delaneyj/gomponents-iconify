package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartAddFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Zm1.243-9.243a6 6 0 0 1 .507 7.91a6 6 0 0 0-8.06 8.127l-.69.691l-8.479-8.492a6 6 0 0 1 8.48-8.464a5.998 5.998 0 0 1 8.242.228Z"/>`),
		g.Group(children),
	)
}