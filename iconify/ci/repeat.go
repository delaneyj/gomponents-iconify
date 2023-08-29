package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 22l-4-4l4-4v3h10v-4h2v5a1 1 0 0 1-1 1H7v3Zm0-11H5V6a1 1 0 0 1 1-1h11V2l4 4l-4 4V7H7v4Z"/>`),
		g.Group(children),
	)
}