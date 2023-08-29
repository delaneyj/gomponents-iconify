package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#FFF" stroke="#000" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="6" d="M63.65 8.07c42.66.21 59.94 22.46 60.26 50.15c.18 15.48-9.79 28.84-25.07 36.98c-7.45 5.27 1.42 22.54 8.52 26.95c0 0-25.1-1.71-32.02-19.4c-3.82.56-7.75.85-11.78.85c-33.28 0-60.17-17.68-60.17-45.38S20.58 7.86 63.65 8.07z"/>`),
		g.Group(children),
	)
}