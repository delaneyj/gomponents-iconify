package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebSection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2h2v20h-2z" opacity=".25"/><path fill="currentColor" d="M3 2h11v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z" opacity=".5"/><path fill="currentColor" d="M16 2h5a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1h-5V2z"/>`),
		g.Group(children),
	)
}