package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dismiss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2c4.42 0 8 3.58 8 8s-3.58 8-8 8s-8-3.58-8-8s3.58-8 8-8zm5 11l-3-3l3-3l-2-2l-3 3l-3-3l-2 2l3 3l-3 3l2 2l3-3l3 3z"/>`),
		g.Group(children),
	)
}