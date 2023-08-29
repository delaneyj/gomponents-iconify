package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockedWithKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#fluentEmojiHighContrastLockedWithKey0)"><path d="M16.5 17.5c0 .818-.393 1.544-1 2V23a1.5 1.5 0 0 1-3 0v-3.5a2.5 2.5 0 1 1 4-2ZM26.5 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M8 6a6 6 0 0 1 12 0v.03A6.5 6.5 0 0 1 25.5 3v.002A6.5 6.5 0 0 1 29 14.98v7.521a3.5 3.5 0 0 1-3 3.464V26a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V13a4 4 0 0 1 4-4h2V6Zm16 19.664a3.498 3.498 0 0 1-2-3.162v-1.754c0-.539.157-1.065.452-1.516a2.713 2.713 0 0 1-.144-.869A2.723 2.723 0 0 0 22 15.978v-1A6.5 6.5 0 0 1 19.173 11H6a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-.336ZM17.5 6a3.5 3.5 0 1 0-7 0v3h7V6Zm8-1.001a4.5 4.5 0 0 0-1.5 8.744v2.232a.735.735 0 0 0 .215.52l.314.31a.723.723 0 0 1 0 1.04a.713.713 0 0 0 0 1.03a.723.723 0 0 1 0 1.04l-.314.31a.716.716 0 0 0-.215.52v1.754a1.5 1.5 0 1 0 3 0v-8.756a4.5 4.5 0 0 0-1.5-8.744Z"/></g><defs><clipPath id="fluentEmojiHighContrastLockedWithKey0"><path fill="#fff" d="M0 0h32v32H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}