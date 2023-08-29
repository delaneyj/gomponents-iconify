package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1.64 12.703C3.502 9.756 8.137 4 16 4c7.862 0 12.497 5.756 14.36 8.703a4.268 4.268 0 0 1 0 4.594c-.868 1.373-2.305 3.307-4.36 5.005v4.2a1.5 1.5 0 0 1-2.184 1.335l-4.35-2.23c-1.08.25-2.236.393-3.466.393c-7.862 0-12.497-5.756-14.36-8.703a4.268 4.268 0 0 1 0-4.594ZM16 6C9.21 6 5.097 10.978 3.33 13.772a2.268 2.268 0 0 0 0 2.456C5.097 19.022 9.21 24 16 24c1.2 0 2.311-.154 3.337-.422l.37-.096L24 25.683v-4.349l.384-.3c2.022-1.58 3.44-3.468 4.286-4.806a2.268 2.268 0 0 0 0-2.456C26.903 10.978 22.79 6 16 6Z"/>`),
		g.Group(children),
	)
}