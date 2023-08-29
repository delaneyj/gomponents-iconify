package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 30H6a2 2 0 0 1-2-2V16a2 2 0 0 1 2-2h3v2H6v12h20V16h-3v-2h3a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2Z"/><path fill="currentColor" d="M13 20h6v2h-6zm7.59-11.41L17 12.17V2h-2v10.17l-3.59-3.58L10 10l6 6l6-6l-1.41-1.41z"/>`),
		g.Group(children),
	)
}