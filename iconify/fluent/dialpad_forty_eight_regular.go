package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 8.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm10.5-20a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM37 8.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 10a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`),
		g.Group(children),
	)
}