package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m25.545 23.328l-7.627-7.705l7.616-7.616l1.857 1.857l2.26-8.428l-8.428 2.258l1.836 1.836l-7.603 7.604l-7.513-7.59L9.81 3.695L1.392 1.394l2.215 8.44l1.848-1.83l7.524 7.604l-7.515 7.515l-1.856-1.855l-2.26 8.427l8.43-2.257L7.94 25.6l7.503-7.502l7.614 7.693l-1.867 1.848l8.416 2.3l-2.213-8.438z"/>`),
		g.Group(children),
	)
}