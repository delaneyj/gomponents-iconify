package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 4v22.75l-7.1-3.59l-.9-.45l-.9.45L8 26.75V4h16m0-2H8a2 2 0 0 0-2 2v26l10-5l10 5V4a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}