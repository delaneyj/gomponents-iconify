package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubyalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896L0 320L256 0h512l256 320zM736 64H288L96 320l416 480l416-480zM512 704V128h190l130 192z"/>`),
		g.Group(children),
	)
}