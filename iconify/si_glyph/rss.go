package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 3c7.355 0 13 5.593 13 12.968h3C17 7.198 9.747 0 1 0v3z"/><path d="M1.052 8.99c4.008 0 6.957 2.9 6.957 6.919h2.95c0-5.346-4.578-9.847-9.907-9.847V8.99zm.01 6.91h3.91c0-2.491-1.43-3.797-3.91-3.82c-.014 0 0 3.82 0 3.82z"/></g>`),
		g.Group(children),
	)
}