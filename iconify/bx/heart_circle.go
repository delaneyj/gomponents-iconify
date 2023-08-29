package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.279 8.833L12 9.112l-.279-.279a2.745 2.745 0 0 0-3.906 0a2.745 2.745 0 0 0 0 3.907L12 16.926l4.186-4.186a2.745 2.745 0 0 0 0-3.907a2.746 2.746 0 0 0-3.907 0z"/><path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/>`),
		g.Group(children),
	)
}