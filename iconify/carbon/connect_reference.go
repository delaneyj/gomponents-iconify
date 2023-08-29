package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4z"/><path fill="currentColor" d="M23 16c-1.57 0-3.012.516-4.181 1.392L14.6 13.175A6.922 6.922 0 0 0 16 9c0-3.86-3.14-7-7-7S2 5.14 2 9s3.14 7 7 7c1.57 0 3.018-.535 4.187-1.411l4.218 4.217A6.986 6.986 0 0 0 16 23c0 3.86 3.14 7 7 7s7-3.14 7-7s-3.14-7-7-7ZM4 9c0-2.757 2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5Z"/>`),
		g.Group(children),
	)
}