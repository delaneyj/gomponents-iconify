package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M15 7a1 1 0 0 0-1 1v9.75a.25.25 0 0 1-.25.25h-3.063c-.9 0-1.342 1.096-.693 1.72l5.313 5.113a1 1 0 0 0 1.386 0l5.313-5.112c.649-.625.207-1.721-.693-1.721H18.25a.25.25 0 0 1-.25-.25V8a1 1 0 0 0-1-1h-2Z"/></g>`),
		g.Group(children),
	)
}