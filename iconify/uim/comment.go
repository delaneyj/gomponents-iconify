package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22H3a1 1 0 0 1-.707-1.707l1.964-1.964A10 10 0 1 1 12 22Z" opacity=".5"/>`),
		g.Group(children),
	)
}