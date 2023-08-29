package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.018 9.5h13.964a7 7 0 0 0-13.964 0ZM2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm14.982.5H3.018a7 7 0 0 0 13.964 0Z"/>`),
		g.Group(children),
	)
}