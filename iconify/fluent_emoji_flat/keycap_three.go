package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M14.291 12.225a1.556 1.556 0 1 1 2.015 2.105a1.75 1.75 0 0 0-.252 3.443l.086.023a1.558 1.558 0 0 1-.446 3.047a1.557 1.557 0 0 1-1.483-1.083a1.75 1.75 0 0 0-3.335 1.061a5.057 5.057 0 1 0 8.737-4.728a5.056 5.056 0 1 0-8.475-5.388a1.75 1.75 0 0 0 3.153 1.52Z"/></g>`),
		g.Group(children),
	)
}