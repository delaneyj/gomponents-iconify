package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridCirclesTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 8A5 5 0 1 0 3 8a5 5 0 0 0 10 0Zm-1.875 0a3.125 3.125 0 1 1-6.25 0a3.125 3.125 0 0 1 6.25 0ZM25 8a5 5 0 1 0-10 0a5 5 0 0 0 10 0Zm-1.875 0a3.125 3.125 0 1 1-6.25 0a3.125 3.125 0 0 1 6.25 0ZM8 25a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0-1.875a3.125 3.125 0 1 0 0-6.25a3.125 3.125 0 0 0 0 6.25ZM25 20a5 5 0 1 0-10 0a5 5 0 0 0 10 0Zm-1.875 0a3.125 3.125 0 1 1-6.25 0a3.125 3.125 0 0 1 6.25 0Z"/>`),
		g.Group(children),
	)
}