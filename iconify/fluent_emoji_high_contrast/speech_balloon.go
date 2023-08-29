package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeechBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.36 12.703C28.498 9.756 23.863 4 16 4C8.138 4 3.503 9.756 1.64 12.703a4.268 4.268 0 0 0 0 4.594c.868 1.373 2.305 3.307 4.36 5.005v4.2a1.5 1.5 0 0 0 2.184 1.335l4.35-2.23A15.27 15.27 0 0 0 16 26c7.862 0 12.497-5.756 14.36-8.703a4.268 4.268 0 0 0 0-4.594ZM16 6c6.79 0 10.903 4.978 12.67 7.772a2.268 2.268 0 0 1 0 2.456C26.903 19.022 22.79 24 16 24c-1.2 0-2.311-.154-3.337-.422l-.37-.096L8 25.683v-4.349l-.384-.3c-2.022-1.58-3.44-3.468-4.286-4.806a2.268 2.268 0 0 1 0-2.456C5.097 10.978 9.21 6 16 6Z"/>`),
		g.Group(children),
	)
}