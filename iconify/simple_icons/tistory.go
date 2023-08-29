package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tistory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 3a3 3 0 1 0 6 0a3 3 0 0 0-6 0m9 18a3 3 0 1 0 6 0a3 3 0 0 0-6 0m0-9a3 3 0 1 0 6 0a3 3 0 0 0-6 0m0-9a3 3 0 1 0 6 0a3 3 0 0 0-6 0m9 0a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/>`),
		g.Group(children),
	)
}