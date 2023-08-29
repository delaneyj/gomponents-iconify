package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoicemailTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 10a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0ZM2 6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6Zm7 4a1.99 1.99 0 0 0-.268-1h2.536A2 2 0 1 0 13 8H7a2 2 0 1 0 2 2Z"/>`),
		g.Group(children),
	)
}