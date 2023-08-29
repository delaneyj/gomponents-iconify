package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PullRequest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17.184V8.816a3 3 0 1 0-2 0v8.368a3 3 0 1 0 2 0Zm14 0V7a2 2 0 0 0-2-2h-5.172l1.586-1.586L13 2l-2.586 2.586L9 6l1.414 1.414L13 10l1.414-1.414L12.828 7H18v10.184a3 3 0 1 0 2 0Z"/>`),
		g.Group(children),
	)
}