package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 11.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm10.5-20a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm10.5-30a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}