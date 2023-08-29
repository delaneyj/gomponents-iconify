package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chopsticks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M253.8 19.46h-1.1c-1.8.1-3.5.34-5.1.85c-4.2 1.3-7.4 3.84-9.6 9.49l-14.5 39.6l33.8 12.74l16.4-38.93c2.2-5.65 1.6-9.64-.6-13.3c-2.2-3.8-6.3-7.04-11.2-8.87c-2.6-1-5.4-1.53-8.1-1.59zM217.4 86.2L82.21 455.4l15.84 5.9L250.4 98.63zm189.8 11.6c-4.3.2-8.1 1.85-11.8 6.7l-24.5 34.6l29.3 20.8l26.1-33.4c3.7-4.8 4.1-8.9 3-12.9c-1.2-4.3-4.2-8.4-8.5-11.4c-3.7-2.67-8-4.21-11.9-4.4h-1.7zm-46.7 55.9L132.1 476l13.7 9.7l243.3-311.6z"/>`),
		g.Group(children),
	)
}