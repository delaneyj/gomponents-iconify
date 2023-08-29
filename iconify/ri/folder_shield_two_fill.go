package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderShieldTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 10H12v7.382c0 1.409.632 2.734 1.705 3.618H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h7.414l2 2H21a1 1 0 0 1 1 1v4Zm-8 2h8v5.382c0 .897-.446 1.734-1.188 2.23L18 21.499l-2.813-1.885A2.684 2.684 0 0 1 14 17.383V12Z"/>`),
		g.Group(children),
	)
}