package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.926 4.121c0 1.68-1.309 4.891-2.922 4.891c-1.613 0-2.923-3.211-2.923-4.891c0-1.68 1.31-3.04 2.923-3.04s2.922 1.36 2.922 3.04zm-1.893 8.675l2.393-2.421l1.279 1.335l1.643-1.662c-.631-.65-1.405-.998-2.669-.998c-.854 1.156-3.689 1.453-3.689 1.453s-2.899-.285-3.753-1.427c-4.093 0-4.217 5.91-4.217 5.91h11.1l-2.087-2.19z"/><path d="m15.094 10.801l-2.381 2.381l-1.416-1.417l-.925.924l2.342 2.341l3.306-3.304l-.926-.925z"/></g>`),
		g.Group(children),
	)
}