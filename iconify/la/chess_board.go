package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessBoard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 4v3h3V4H7zm3 3v3h3V7h-3zm3 0h3V4h-3v3zm3 0v3h3V7h-3zm3 0h3V4h-3v3zm3 0v3h3V7h-3zm3 0h3V4h-3v3zm0 3v3h3v-3h-3zm0 3h-3v3h3v-3zm0 3v3h3v-3h-3zm0 3h-3v3h3v-3zm0 3v3h3v-3h-3zm0 3h-3v3h3v-3zm-3 0v-3h-3v3h3zm-3 0h-3v3h3v-3zm-3 0v-3h-3v3h3zm-3 0h-3v3h3v-3zm-3 0v-3H7v3h3zm-3 0H4v3h3v-3zm0-3v-3H4v3h3zm0-3h3v-3H7v3zm0-3v-3H4v3h3zm0-3h3v-3H7v3zm0-3V7H4v3h3zm3 3v3h3v-3h-3zm3 0h3v-3h-3v3zm3 0v3h3v-3h-3zm3 0h3v-3h-3v3zm0 3v3h3v-3h-3zm0 3h-3v3h3v-3zm-3 0v-3h-3v3h3zm-3 0h-3v3h3v-3z"/>`),
		g.Group(children),
	)
}