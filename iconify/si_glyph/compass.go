package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.021.035c-4.411 0-8 3.588-8 8c0 4.413 3.588 8 8 8c4.411 0 8-3.587 8-8c0-4.412-3.589-8-8-8zM9.001 14A6.006 6.006 0 0 1 3 8c0-3.307 2.692-6 6-6c3.31 0 6 2.693 6 6c0 3.308-2.69 6-6 6z"/><path d="m6.042 6.021l2.021 3L12.98 11l-2.979-3.958l-3.959-1.021z"/></g>`),
		g.Group(children),
	)
}