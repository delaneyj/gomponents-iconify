package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cyclist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 30a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6zm0-10a4 4 0 1 0 4 4a4.005 4.005 0 0 0-4-4zM7 30a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6zm0-10a4 4 0 1 0 4 4a4.005 4.005 0 0 0-4-4z"/><path fill="currentColor" d="M17 27h-2v-6.586L9.585 15a2.003 2.003 0 0 1 0-2.83l4.586-4.585a2.002 2.002 0 0 1 2.828 0L21.414 12H27v2h-6.415l-5-5L11 13.585l6 6zm4.5-19A3.5 3.5 0 1 1 25 4.5A3.504 3.504 0 0 1 21.5 8zm0-5A1.5 1.5 0 1 0 23 4.5A1.502 1.502 0 0 0 21.5 3z"/>`),
		g.Group(children),
	)
}