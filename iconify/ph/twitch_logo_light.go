package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitchLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 34H48a14 14 0 0 0-14 14v144a14 14 0 0 0 14 14h18v34a6 6 0 0 0 9.84 4.61l45.78-38.15a2 2 0 0 1 1.28-.46h42.2a14 14 0 0 0 9-3.25L217 167a14 14 0 0 0 5-10.76V48a14 14 0 0 0-14-14Zm2 122.25a2 2 0 0 1-.72 1.54l-42.9 35.75a2 2 0 0 1-1.28.46h-42.2a14 14 0 0 0-9 3.25L78 227.19V200a6 6 0 0 0-6-6H48a2 2 0 0 1-2-2V48a2 2 0 0 1 2-2h160a2 2 0 0 1 2 2ZM174 88v48a6 6 0 0 1-12 0V88a6 6 0 0 1 12 0Zm-48 0v48a6 6 0 0 1-12 0V88a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}