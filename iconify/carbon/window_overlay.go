package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowOverlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M15 6h2v3h-2z" fill="currentColor"/><path d="M25 17h3v2h-3z" fill="currentColor"/><path d="M15 27h2v3h-2z" fill="currentColor"/><path d="M4 17h3v2H4z" fill="currentColor"/><path d="M6.774 10.187l1.414-1.415l2.121 2.122l-1.414 1.414z" fill="currentColor"/><path d="M21.694 10.898l2.121-2.122l1.414 1.415l-2.12 2.12z" fill="currentColor"/><path d="M21.69 25.087l1.414-1.415l2.121 2.122l-1.414 1.414z" fill="currentColor"/><path d="M6.777 25.81l2.121-2.12l1.414 1.414l-2.121 2.12z" fill="currentColor"/><path d="M4 2h24v2H4z" fill="currentColor"/><path d="M16 24a6 6 0 1 0-6-6a6 6 0 0 0 6 6zm0-10v8a4 4 0 0 1 0-8z" fill="currentColor"/>`),
		g.Group(children),
	)
}