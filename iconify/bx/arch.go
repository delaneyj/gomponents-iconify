package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22h1v-6.995c.006-.502.177-3.005 3-3.005s2.994 2.503 3 3v7h7v-2h-1V4h1V2H2v2h1v16H2v2h6zM19 4v2H5V4h14zM5 8h14v12h-2v-5c0-1.729-1.045-5-5-5s-5 3.271-5 5v5H5V8z"/>`),
		g.Group(children),
	)
}