package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 9a6 6 0 1 1 9.642 4.769C14.186 14.907 16 17.208 16 20a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1c0-2.792 1.814-5.093 4.358-6.231A5.99 5.99 0 0 1 2 9zm15 2a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2h-6z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}