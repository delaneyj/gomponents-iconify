package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 5.4A2.4 2.4 0 0 1 4.4 3h.1a.5.5 0 0 1 0 1h-.1A1.4 1.4 0 0 0 3 5.4v.2A1.4 1.4 0 0 0 4.4 7h.1a.5.5 0 0 1 0 1h-.1A2.4 2.4 0 0 1 2 5.6v-.2Zm8 0A2.4 2.4 0 0 0 7.6 3h-.1a.5.5 0 0 0 0 1h.1A1.4 1.4 0 0 1 9 5.4v.2A1.4 1.4 0 0 1 7.6 7h-.1a.5.5 0 0 0 0 1h.1A2.4 2.4 0 0 0 10 5.6v-.2ZM4.5 5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Z"/>`),
		g.Group(children),
	)
}