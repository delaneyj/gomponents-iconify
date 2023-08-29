package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 22l-3-3H5a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2h-4l-3 3ZM5 4v13h5l2 2l2-2h5V4H5Zm11 10.986H8V14.6a3.82 3.82 0 0 1 4-3.6a3.82 3.82 0 0 1 4 3.6v.386ZM12 10a1.934 1.934 0 0 1-2-2a1.935 1.935 0 0 1 2-2a1.935 1.935 0 0 1 2 2a1.934 1.934 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}