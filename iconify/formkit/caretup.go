package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caretup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.18 1.38L3.03 6.21c-.6.7-.1 1.79.82 1.79h8.29c.93 0 1.42-1.09.82-1.79L8.82 1.38c-.43-.5-1.21-.5-1.64 0Z"/>`),
		g.Group(children),
	)
}