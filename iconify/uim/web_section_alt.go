package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebSectionAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.9 2H8.1A2.584 2.584 0 0 0 8 3v18c-.032.337.002.676.1 1h1.8c.098-.324.132-.663.1-1V3a2.584 2.584 0 0 0-.1-1z" opacity=".25"/><path fill="currentColor" d="M3 2h5v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M10 2h11a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H10V2z" opacity=".5"/>`),
		g.Group(children),
	)
}