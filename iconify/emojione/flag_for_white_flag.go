package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForWhiteFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d0d0" d="M48.2.7c-16.6-5.2-22.4 19.4-39 14.2c3 8.3 6 16.5 9 24.8c16.6 5.2 22.4-19.4 39-14.2c-3-8.3-6-16.5-9-24.8"/><path fill="#3e4347" d="M6.8 15.8L24.4 64h1L7.7 15.4l.6-.2L26 64h1L9.1 14.9z"/><path fill="#94989b" d="m8.3 15.2l-.6.2L25.4 64h.6z"/>`),
		g.Group(children),
	)
}