package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M10 25.593a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h1.986a.013.013 0 0 0 .014-.014c0-.008.007-.014.015-.013c.156.018.318.027.485.027c2.197 0 3.5-1.599 3.5-3v-5.75a.25.25 0 0 0-.25-.25h-1.988a.75.75 0 0 1-.528-1.283l4.238-4.195a.75.75 0 0 1 1.056 0l4.238 4.195a.75.75 0 0 1-.528 1.283H20.25a.25.25 0 0 0-.25.25v5.75c0 3.866-3.358 7-7.5 7c-.165 0-.33-.005-.492-.015a.007.007 0 0 0-.008.008a.007.007 0 0 1-.007.007H10Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}