package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.75 4c0-.69.56-1.25 1.25-1.25h10c.69 0 1.25.56 1.25 1.25v14a1.25 1.25 0 0 1-2.137.881L11 14.741l-4.113 4.14A1.25 1.25 0 0 1 4.75 18V4Zm2.5 1.25v9.718l2.863-2.882a1.25 1.25 0 0 1 1.774 0l2.863 2.882V5.25h-7.5Z" opacity=".2"/><path d="M4.5 3a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.855.352L10 12.676l-4.645 4.676A.5.5 0 0 1 4.5 17V3Zm1 .5v12.287l4.145-4.172a.5.5 0 0 1 .71 0l4.145 4.172V3.5h-9Z"/></g>`),
		g.Group(children),
	)
}