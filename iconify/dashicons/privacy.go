package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Privacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 9.6c-.6 0-1 .4-1 1c0 .4.3.7.6.8l-.3 1.4h1.3l-.3-1.4c.4-.1.6-.4.6-.8c.1-.6-.3-1-.9-1zm.1-4.3c-.7 0-1.4.5-1.4 1.2V8h2.7V6.5c-.1-.7-.6-1.2-1.3-1.2zM10 2L3 5v3c.1 4.4 2.9 8.3 7 9.9c4.1-1.6 6.9-5.5 7-9.9V5l-7-3zm4 11c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V9c0-.6.4-1 1-1h.3V6.5C7.4 5.1 8.6 4 10 4c1.4 0 2.6 1.1 2.7 2.5V8h.3c.6 0 1 .4 1 1v4z"/>`),
		g.Group(children),
	)
}