package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosscutSaw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M333.2 25.24c-2.6 2.79-4.7 8.5 4.3 17.93l53.2 53.59L406 82.83l-54.2-53.11c-10.5-9.45-16-7.26-18.6-4.48zM25.21 333.3c-2.7 2.6-4.97 8.1 4.48 17.8l53.11 55l15.05-15l-54.65-53.5c-9.49-9-15.22-6.9-17.99-4.3zm72.64 85.5l69.65 70.3l3.5-31.3l26.5 15.9c.1-.1.2-.1.3-.1l5.2-33.7l32.7 8.8l.2-.1l9-36.4l34 7.8c.1-.1.3-.3.5-.4l6.5-38.1l31 6.7l2.8-35.1l35 1.6l-1.4-35l34.9-2.8l-6.7-30.7l38-6.8l.4-.5l-7.7-34.1l36.7-8.8l.2-.2l-9.2-32.7l33.5-6.7l.2-.2l-15.9-25.1l31.4-3.5l-70.4-70.84z"/>`),
		g.Group(children),
	)
}