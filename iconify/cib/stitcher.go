package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stitcher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.974 9.25h-.479v13.5h.479zm-.995 1.5H22.5v10.5h.479zm-1.489 0H0v10.5h21.49zm10.51 0h-6.01v10.5H32zm-7.01 0h-.479v10.5h.479z"/>`),
		g.Group(children),
	)
}