package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SilentSquint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.66 12.21a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l1.5-1.5a1 1 0 0 0 0-1.42l-1.5-1.5a1 1 0 0 0-1.42 1.42l.8.79l-.8.79a1 1 0 0 0 0 1.42Zm7.5 0a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.79-.79l.79-.79a1 1 0 1 0-1.42-1.42l-1.5 1.5a1 1 0 0 0 0 1.42Zm.11 2a1 1 0 0 0-1.2 0l-.74.55l-.73-.55a1 1 0 0 0-1.2 0l-.73.55l-.74-.55a1 1 0 0 0-1.2 0l-1.33 1a1 1 0 1 0 1.2 1.6l.73-.55l.74.55a.67.67 0 0 0 .12.06a.83.83 0 0 0 .22.08h.48a1.12 1.12 0 0 0 .33-.16l.73-.55l.73.55a1 1 0 0 0 1 .11l.1-.05a.39.39 0 0 0 .11-.06l.74-.55l.73.55a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4a1 1 0 0 0-.2-1.4ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}