package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageFacingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#B4ACBC" d="M20.343 2.293A1 1 0 0 0 19.636 2H7a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V9.364a1 1 0 0 0-.293-.707l-6.364-6.364Z"/><path fill="#F3EEF8" d="M19.682 3H7a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V9.453L19.682 3Z"/><path fill="#998EA4" d="M9.5 12h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1 0-1Zm0 3a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1h-13ZM9 18.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5Zm.5 2.5a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-8Z"/><path fill="#CDC4D6" d="M26 9.453h-4.61a1.707 1.707 0 0 1-1.708-1.707V3L26 9.453Z"/></g>`),
		g.Group(children),
	)
}