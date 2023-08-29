package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 6A5 5 0 1 1 1 6a5 5 0 0 1 10 0ZM9.5 6c0-.695-.203-1.342-.552-1.887L4.113 8.948A3.5 3.5 0 0 0 9.5 6ZM7.888 3.052a3.5 3.5 0 0 0-4.836 4.836l4.836-4.836Z"/>`),
		g.Group(children),
	)
}