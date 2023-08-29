package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 64c-15.4 3.77-35.7 16.04-53 33.17c-19.2 19.13-34.9 43.63-39.58 64.63l-.58 135.3l.37 157.4l93.99-40.3L192 64zm128 0l-1.2 350.2l94 40.3l.4-156.8l-.6-135.9c-4.7-21-20.3-45.5-39.6-64.63c-17.3-17.13-37.6-29.4-53-33.17zM57.24 94.67c-8.39 0-15 6.63-15 15.03c0 8.4 6.61 15 15 15s15-6.6 15-15s-6.61-15.03-15-15.03zm397.56 0c-8.4 0-15 6.63-15 15.03c0 8.4 6.6 15 15 15s15-6.6 15-15s-6.6-15.03-15-15.03zM35.5 142.7l-1.42 334h46l1.42-334h-46zm395 0l1.5 334h46l-1.5-334h-46zM159.2 231h18v48h-18v-48zm175.6 0h18v48h-18v-48z"/>`),
		g.Group(children),
	)
}