package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 5C9.937 5 5 9.938 5 16v8c0 1.645 1.355 3 3 3h3v-9H7v-2c0-4.984 4.016-9 9-9s9 4.016 9 9v2h-4v9h3c1.645 0 3-1.355 3-3v-8c0-6.063-4.938-11-11-11zM7 20h2v5H8c-.566 0-1-.434-1-1zm16 0h2v4c0 .566-.434 1-1 1h-1z"/>`),
		g.Group(children),
	)
}