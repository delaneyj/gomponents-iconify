package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 23h8v2H4zm20.523-9H4v2h20.5a3.5 3.5 0 0 1 0 7h-5.672l2.586-2.586L20 19l-5 5l5 5l1.414-1.414L18.828 25h5.705a5.5 5.5 0 0 0-.01-11zM4 5h24v2H4z"/>`),
		g.Group(children),
	)
}