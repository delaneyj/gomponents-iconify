package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moustache(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 9a3 3 0 0 1 2.599 1.5h0c.933 1.333 2.133 1.556 3.126 1.556h.291l.77-.044h.213c-.963 1.926-3.163 2.925-6.6 3h-.565a3 3 0 0 1 .165-6z"/><path d="M9 9a3 3 0 0 0-2.599 1.5h0c-.933 1.333-2.133 1.556-3.126 1.556h-.291l-.77-.044h-.213c.963 1.926 3.163 2.925 6.6 3h.565a3 3 0 0 0-.165-6z"/></g>`),
		g.Group(children),
	)
}