package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digitalocean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 6a6 6 0 0 0-6 6H1C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11v-5a6 6 0 0 0 0-12Z"/><path d="M7 18v-5h5v5H7Zm-4 0v4h4v-4H3Zm0 0H1v-2h2v2Z"/></g>`),
		g.Group(children),
	)
}