package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsMessageAltX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M8.5 18l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h3.5zM7.293 6.707l1.414-1.414L12 8.586l3.293-3.293l1.414 1.414L13.414 10l3.293 3.293l-1.414 1.414L12 11.414l-3.293 3.293l-1.414-1.414L10.586 10L7.293 6.707z" fill="currentColor"/>`),
		g.Group(children),
	)
}