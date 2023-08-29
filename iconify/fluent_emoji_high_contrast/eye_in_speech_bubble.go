package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeInSpeechBubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16 22a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M2.484 16.237C4.3 13.367 8.674 8 16 8s11.7 5.367 13.515 8.237a3.268 3.268 0 0 1 0 3.526c-.883 1.397-2.373 3.386-4.515 5.06v4.679a.5.5 0 0 1-.728.445l-4.683-2.401A14.17 14.17 0 0 1 16 28c-7.326 0-11.7-5.367-13.516-8.237a3.267 3.267 0 0 1 0-3.526ZM23 18a7 7 0 1 0-14 0a7 7 0 0 0 14 0Z"/></g>`),
		g.Group(children),
	)
}