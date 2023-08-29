package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M28.286 37h11.428M42 42l-2.286-5L42 42Zm-16 0l2.286-5L26 42Zm2.286-5L34 24l5.714 13H28.286ZM16 6l1 3M6 11h22m-18 5s1.79 6.26 6.263 9.74C20.737 29.216 28 32 28 32"/><path d="M24 11s-1.79 8.217-6.263 12.783C13.263 28.348 6 32 6 32"/></g>`),
		g.Group(children),
	)
}