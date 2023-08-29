package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.5 11c-.706 0-1.342.293-1.797.763L4.969 8.396a2.46 2.46 0 0 0 0-.792l6.734-3.367a2.5 2.5 0 1 0-.672-1.341L4.297 6.263a2.5 2.5 0 1 0 0 3.474l6.734 3.367A2.5 2.5 0 1 0 13.5 11z"/>`),
		g.Group(children),
	)
}