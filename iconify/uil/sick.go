package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm6 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm.27 3.2a1 1 0 0 0-1.2 0l-.74.55l-.73-.55a1 1 0 0 0-1.2 0l-.73.55l-.74-.55a1 1 0 0 0-1.2 0l-1.33 1a1 1 0 1 0 1.2 1.6l.73-.55l.74.55l.12.06a.83.83 0 0 0 .22.08h.12a1 1 0 0 0 .25 0h.1a1.06 1.06 0 0 0 .34-.16l.73-.55l.73.55a1 1 0 0 0 1 .11l.1-.05a.39.39 0 0 0 .11-.06l.74-.55l.73.55a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4a1 1 0 0 0-.2-1.4ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}