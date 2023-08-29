package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M36 32a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28z"/><path fill="#FFF" d="M26.621 9.001c0-1.24-.93-2.263-2.232-2.263c-.807 0-1.52.341-1.922.93l-4.465 6.48l-4.465-6.48c-.403-.589-1.116-.93-1.922-.93c-1.302 0-2.232 1.023-2.232 2.263c0 .496.155.992.434 1.396L15 17.692v8.417A2.89 2.89 0 0 0 17.891 29h.219A2.89 2.89 0 0 0 21 26.109v-8.414l5.188-7.299c.279-.403.433-.899.433-1.395z"/>`),
		g.Group(children),
	)
}