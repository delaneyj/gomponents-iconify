package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosMicOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M256 32c-43.7 0-79 37.5-79 83.5V270c0 46 35.3 83.5 79 83.5s79-37.5 79-83.5V115.5c0-46-35.3-83.5-79-83.5zm63 238c0 37.2-28.3 67.5-63 67.5s-63-30.3-63-67.5V115.5c0-37.2 28.3-67.5 63-67.5s63 30.3 63 67.5V270z" fill="currentColor"/><path d="M367 192v79.7c0 60.2-49.8 109.2-110 109.2s-110-49-110-109.2V192h-19v79.7c0 67.2 53 122.6 120 127.5V462h-73v18h161v-18h-69v-62.8c66-4.9 117-60.3 117-127.5V192h-17z" fill="currentColor"/>`),
		g.Group(children),
	)
}