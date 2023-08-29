package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M20 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4ZM9 24a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H9Zm2 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm4-1a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1ZM7 11a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/><path d="M20 5a6.001 6.001 0 0 0-5.5 3.598a6.001 6.001 0 1 0-8.528 7.583A3.001 3.001 0 0 0 4 19v1H3v-.5a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0V23h1v5a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-.118a2 2 0 0 0 1.106 1.789l3 1.5A2 2 0 0 0 31 29.382V17.618a2 2 0 0 0-2.894-1.789l-3 1.5A2 2 0 0 0 24 19.12V19c0-1.089-.58-2.042-1.448-2.568A6 6 0 0 0 20 5Zm-4 6a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm13 6.618v11.764l-3-1.5v-8.764l3-1.5ZM7 18h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1ZM9 7a4 4 0 1 1 0 8a4 4 0 0 1 0-8Z"/></g>`),
		g.Group(children),
	)
}