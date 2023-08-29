package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 897H256L0 449L256 0h512l256 449zM736 64H288L64 449l224 384h448l224-384z"/>`),
		g.Group(children),
	)
}