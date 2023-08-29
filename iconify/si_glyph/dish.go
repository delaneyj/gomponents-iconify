package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.979 4.06a3.933 3.933 0 0 0-3.926 3.929a3.933 3.933 0 0 0 3.926 3.931a3.933 3.933 0 0 0 3.926-3.931A3.933 3.933 0 0 0 8.979 4.06z"/><path d="M9.007.072c-4.399 0-7.964 3.562-7.964 7.957a7.96 7.96 0 0 0 7.964 7.96c4.397 0 7.965-3.562 7.965-7.96c0-4.394-3.568-7.957-7.965-7.957zm-.028 13.064a5.15 5.15 0 0 1-5.143-5.147a5.15 5.15 0 0 1 5.143-5.146a5.15 5.15 0 0 1 5.143 5.146a5.152 5.152 0 0 1-5.143 5.147z"/></g>`),
		g.Group(children),
	)
}