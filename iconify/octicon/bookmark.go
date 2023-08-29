package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M9 0H1C.27 0 0 .27 0 1v15l5-3.09L10 16V1c0-.73-.27-1-1-1zm-.78 4.25L6.36 5.61l.72 2.16c.06.22-.02.28-.2.17L5 6.6L3.12 7.94c-.19.11-.25.05-.2-.17l.72-2.16l-1.86-1.36c-.17-.16-.14-.23.09-.23l2.3-.03l.7-2.16h.25l.7 2.16l2.3.03c.23 0 .27.08.09.23h.01z" fill="currentColor"/>`),
		g.Group(children),
	)
}